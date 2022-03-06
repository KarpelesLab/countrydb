#!/pkg/main/dev-lang.php.core.cli/bin/php
<?php
// load intl if needed
if (!extension_loaded('intl')) dl('intl');

if (!file_exists('iso-codes')) {
	system('git clone https://salsa.debian.org/iso-codes-team/iso-codes.git');
}

// read lists
$country_data = json_decode(file_get_contents('iso-codes/data/iso_3166-1.json'), true)['3166-1'];
$region_data = json_decode(file_get_contents('iso-codes/data/iso_3166-2.json'), true)['3166-2'];

echo 'Loaded '.count($country_data).' countries and '.count($region_data).' regions'."\n";

// compute unique names for countries based on "common_name" or "name" (or special if needed)
$unique_names = [];
$alpha2_index = [];
$special = [
	'CD' => 'DemocraticCongo',
	'VG' => 'BritishVirginIslands',
];

foreach($country_data as &$country) {
	if (isset($special[$country['alpha_2']])) {
		$name = $special[$country['alpha_2']];
	} else if (isset($country['common_name'])) {
		$name = preg_replace('/[^a-zA-Z]/', '', asciify($country['common_name']));
	} else {
		$name = $country['name'];
		list($name) = explode(',', $name);
		$name = preg_replace('/[^a-zA-Z]/', '', asciify($name));
	}
	if (isset($unique_names[strtolower($name)])) die("name not unique: $name\n");
	$unique_names[strtolower($name)] = true;
	$country['unique_name'] = $name;
	$alpha2_index[$country['alpha_2']] = &$country;
}

// ccTLD filling
foreach($country_data as &$country) {
	$cctld = '.'.strtolower($country['alpha_2']); // general rule
	switch($country['alpha_2']) {
	case 'GB': $cctld = '.uk'; break; // UK exception
	case 'BV': $cctld = null; break; // Bouvet Island doesn't have an enabled cctld but is part of Norway
	case 'EH': $cctld = null; break;
	case 'UM': $cctld = null; break;
	}
	$country['cctld'] = $cctld;
}

// fips 10 filling
$fp = fopen('data/fips-10-4-to-iso-country-codes.csv', 'r');
while(!feof($fp)) {
	$lin = fgetcsv($fp);
	// "FIPS 10-4","ISO 3166",Name
	if ($lin === false) break;
	if (strlen($lin[0]) != 2) continue;
	if (strlen($lin[1]) != 2) continue;
	$alpha2_index[$lin[1]]['fips'] = $lin[0];
}

// generate files
foreach($country_data as &$country) {
	$fp = fopen(strtolower('data-'.strtolower($country['unique_name']).'.go'), 'w');
	fwrite($fp, "package countrydb\n\n");
	fwrite($fp, "var ".$country['unique_name']." = &Country{\n");
	fwrite($fp, "\tName: ".goescape($country['name']).",\n");
	fwrite($fp, "\tUniqueName: ".goescape($country['unique_name']).",\n");
	if (isset($country['official_name'])) fwrite($fp, "\tOfficialName: ".goescape($country["official_name"]).",\n");
	if (isset($country['common_name'])) fwrite($fp, "\tCommonName: ".goescape($country["common_name"]).",\n");
	fwrite($fp, "\tISO3166_Alpha2: ".goescape($country['alpha_2']).",\n");
	fwrite($fp, "\tISO3166_Alpha3: ".goescape($country['alpha_3']).",\n");
	fwrite($fp, "\tNumeric: ".((int)$country['numeric']).",\n");
	fwrite($fp, "\tCcTLD: ".goescape($country['cctld']).",\n");
	if (isset($country['fips'])) fwrite($fp, "\tFIPS: ".goescape($country['fips']).",\n");

	fwrite($fp, "}\n");
	fclose($fp);
}

// generate index
$indices = [
	'ByAlpha2' => 'alpha_2',
	'ByAlpha3' => 'alpha_3',
	'ByNumeric' => 'numeric',
	'ByUniqueName' => 'unique_name',
	'ByCcTLD' => 'cctld',
	'ByFIPS' => 'fips',
];
foreach($indices as $name => $col) {
	$fp = fopen('index-'.strtolower($name).'.go', 'w');
	fwrite($fp, "package countrydb\n\n");
	$type = 'string';
	if ($col == 'numeric') $type = 'int';
	fwrite($fp, "var $name = map[$type]*Country{\n");
	foreach($country_data as &$country) {
		$k = $country[$col]??null;
		if (is_null($k)) continue;
		if ($col == 'numeric') {
			$k = (int)$k;
		} else {
			$k = goescape($k);
		}
		fwrite($fp, "\t$k: ".$country['unique_name'].",\n");
	}

	fwrite($fp, "}\n");
	fclose($fp);
}

function asciify($var) {
	return transliterator_transliterate('Latin-ASCII', transliterator_transliterate('Any-Latin', $var));
}

function goescape($str) {
	if (is_null($str)) return '""';
	return '"'.addcslashes($str, "\0..\37\"").'"';
}
