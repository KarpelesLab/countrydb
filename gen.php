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

	fwrite($fp, "}\n");
	fclose($fp);
}

// generate index
$indices = [
	'ByAlpha2' => 'alpha_2',
	'ByAlpha3' => 'alpha_3',
	'ByNumeric' => 'numeric',
	'ByUniqueName' => 'unique_name',
];
foreach($indices as $name => $col) {
	$fp = fopen('index-'.strtolower($name).'.go', 'w');
	fwrite($fp, "package countrydb\n\n");
	$type = 'string';
	if ($col == 'numeric') $type = 'int';
	fwrite($fp, "var $name = map[$type]*Country{\n");
	foreach($country_data as &$country) {
		$k = $country[$col];
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
	return '"'.addcslashes($str, "\0..\37\"").'"';
}
