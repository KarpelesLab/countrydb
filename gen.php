#!/pkg/main/dev-lang.php.core.cli/bin/php
<?php
// load intl if needed
if (!extension_loaded('intl')) dl('intl');

if (!file_exists('iso-codes')) {
	system('git clone https://salsa.debian.org/iso-codes-team/iso-codes.git');
}

$gen = new CountryGenerator();
$gen->process();

class CountryGenerator {
	private $country_data;
	private $region_data;
	private $alpha2_index = [];
	private $golng;
	private $lngIndex = [];

	public function process() {
		$this->loadIsoData();
		$this->loadCctld();
		$this->loadFips10();
		$this->loadCurrency();
		$this->loadPhone();
		$this->loadIcann();
		$this->loadGolng();

		$this->genCountryData();
		$this->genCountryIndex();
		$this->genAll();

		$this->genLanguagesData();
		$this->genLanguagesIndex();
	}

	public function loadIsoData() {
		// read lists
		$this->country_data = json_decode(file_get_contents('iso-codes/data/iso_3166-1.json'), true)['3166-1'];
		$this->region_data = json_decode(file_get_contents('iso-codes/data/iso_3166-2.json'), true)['3166-2'];

		echo 'Loaded '.count($this->country_data).' countries and '.count($this->region_data).' regions'."\n";

		// compute unique names for countries based on "common_name" or "name" (or special if needed)
		$unique_names = [];
		$special = [
			'CD' => 'DemocraticCongo',
			'VG' => 'BritishVirginIslands',
			'TR' => 'Turkey',
		];

		foreach($this->country_data as &$country) {
			if (isset($special[$country['alpha_2']])) {
				$name = $special[$country['alpha_2']];
			} else if (isset($country['common_name'])) {
				$name = preg_replace('/[^a-zA-Z]/', '', self::asciify($country['common_name']));
			} else {
				$name = $country['name'];
				list($name) = explode(',', $name);
				$name = preg_replace('/[^a-zA-Z]/', '', self::asciify($name));
			}
			if ($country['alpha_2'] == 'TW') {
				// Fix TW's name
				$country['official_name'] = $country['common_name'];
				$country['name'] = $country['common_name'];
			}
			if (isset($unique_names[strtolower($name)])) die("name not unique: $name\n");
			$unique_names[strtolower($name)] = true;
			$country['unique_name'] = $name;
			$this->alpha2_index[$country['alpha_2']] = &$country;
		}
	}

	public function loadCctld() {
		// ccTLD filling
		foreach($this->country_data as &$country) {
			$cctld = '.'.strtolower($country['alpha_2']); // general rule
			switch($country['alpha_2']) {
			case 'GB': $cctld = '.uk'; break; // UK exception
			case 'BV': $cctld = null; break; // Bouvet Island doesn't have an enabled cctld but is part of Norway
			case 'EH': $cctld = null; break;
			case 'UM': $cctld = null; break;
			}
			$country['cctld'] = $cctld;
		}
	}

	public function loadFips10() {
		// fips 10-4 filling
		$fp = fopen('data/fips-10-4-to-iso-country-codes.csv', 'r');
		while(!feof($fp)) {
			$lin = fgetcsv($fp);
			// "FIPS 10-4","ISO 3166",Name
			if ($lin === false) break;
			if (strlen($lin[0]) != 2) continue;
			if (strlen($lin[1]) != 2) continue;
			$this->alpha2_index[$lin[1]]['fips'] = $lin[0];
		}
		fclose($fp);
	}

	public function loadCurrency() {
		// currency filling
		$fp = fopen('data/country-to-currency.csv', 'r');
		while(!feof($fp)) {
			$lin = fgetcsv($fp);
			if ($lin === false) break;
			if (strlen($lin[0]) != 2) continue;
			if (strlen($lin[1]) != 3) continue;
			$this->alpha2_index[$lin[0]]['currency'] = $lin[1];
		}
		fclose($fp);
	}

	public function loadPhone() {
		// phone codes filling
		$fp = fopen('data/phone_codes.csv', 'r');
		while(!feof($fp)) {
			$lin = fgetcsv($fp);
			if ($lin === false) break;
			if (strlen($lin[0]) != 2) continue;
			$this->alpha2_index[$lin[0]]['phone_prefix'] = $lin[1];
		}
		fclose($fp);
	}

	public function loadIcann() {
		// icann regions
		$icann_regions = [];
		$fp = fopen('data/icann_regions.csv', 'r');
		while(!feof($fp)) {
			$lin = fgetcsv($fp);
			if ($lin === false) break;
			$icann_regions[$lin[1]] = $lin[0];
		}
		foreach($this->country_data as &$country) {
			if (isset($icann_regions[$country['cctld']]))
				$country['icann_region'] = $icann_regions[$country['cctld']];
		}
	}

	public function loadGolng() {
		// load list of languages known to go
		$fp = fopen('data/golanguages.csv', 'r');
		$this->golng = [];
		while(!feof($fp)) {
			$lin = fgetcsv($fp);
			if ($lin === false) break;
			$this->golng[$lin[1]] = $lin[0]; // en => English
		}
		fclose($fp);
	}

	public function genCountryData() {
		// generate files
		foreach($this->country_data as &$country) {
			$fp = fopen(strtolower('data-'.strtolower($country['unique_name']).'.go'), 'w');
			fwrite($fp, "package countrydb\n\n");
			fwrite($fp, "var ".$country['unique_name']." = &Country{\n");
			fwrite($fp, "\tName: ".self::goescape($country['name']).",\n");
			fwrite($fp, "\tUniqueName: ".self::goescape($country['unique_name']).",\n");
			fwrite($fp, "\tISO3166_Alpha2: ".self::goescape($country['alpha_2']).",\n");
			fwrite($fp, "\tISO3166_Alpha3: ".self::goescape($country['alpha_3']).",\n");
			fwrite($fp, "\tNumeric: ".((int)$country['numeric']).",\n");
			fwrite($fp, "\tCcTLD: ".self::goescape($country['cctld']).",\n");
			if (isset($country['fips'])) fwrite($fp, "\tFIPS: ".self::goescape($country['fips']).",\n");
			if (isset($country['currency'])) fwrite($fp, "\tCurrency: ".self::goescape($country['currency']).",\n");
			if (isset($country['icann_region'])) fwrite($fp, "\tICANN_Region: ".self::goescape($country['icann_region']).",\n");
			if (isset($country['phone_prefix'])) fwrite($fp, "\tPhonePrefix: ".self::goescape($country['phone_prefix']).",\n");

			fwrite($fp, "}\n");
			fclose($fp);
		}
	}

	public function genCountryIndex() {
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
			foreach($this->country_data as &$country) {
				$k = $country[$col]??null;
				if (is_null($k)) continue;
				if ($col == 'numeric') {
					$k = (int)$k;
				} else {
					$k = self::goescape($k);
				}
				fwrite($fp, "\t$k: ".$country['unique_name'].",\n");
			}

			fwrite($fp, "}\n");
			fclose($fp);
		}
	}

	public function genAll() {
		// generate all
		$fp = fopen('all.go', 'w');
		fwrite($fp, "package countrydb\n\n");
		fwrite($fp, "var All = []*Country{\n");
		foreach($this->country_data as &$country) {
			fwrite($fp, "\t".$country['unique_name'].",\n");
		}
		fwrite($fp, "}\n");
		fclose($fp);
	}

	public function genLanguagesData() {
		$this->writeLanguage('en', function($a) { return $a; });

		// generate languages
		foreach(glob('iso-codes/iso_3166-1/*.po') as $file) {
			$lng = basename($file, '.po'); // remove path & .po
			$this->writeLanguage($lng, $this->loadLanguage($file));
		}
	}

	public function loadLanguage($file) {
		// let's read all translations
		$po_data = file_get_contents($file);
		preg_match_all('/\nmsgid "(.+?)"\nmsgstr "(.+?)"/m', $po_data, $trans, PREG_SET_ORDER);

		// index translations
		$trans_idx = [];
		foreach($trans as $t)
			$trans_idx[$t[1]] = stripslashes($t[2]);

		return function($a) use (&$trans_idx) { if (!isset($trans_idx[$a])) return NULL; return $trans_idx[$a]; };
	}

	public function writeLanguage($lng, $translate) {
		// some cases: we may have "ln" "lng" "ln@var" "ln_XX"
		// all of these will parse as language.Tag
		$uniqueName = str_replace('@', '_', $lng);
		if (isset($this->golng[$lng])) {
			$lngVar = $this->golng[$lng];
		} else {
			$lngVar = 'Locale'.str_replace('_', '', $uniqueName); // var containing this language
		}
		$this->lngIndex[$lng] = $lngVar;

		// create file
		$fp = fopen('countrynames/lng-'.strtolower(str_replace('_', '', $uniqueName)).'.go', 'w');
		fwrite($fp, "package countrynames\n\n");
		fwrite($fp, "var $lngVar = map[string]Translated{\n");

		// analyse which country has which translation
		$vars = [
			'name' => 'Name',
			'official_name' => 'OfficialName',
			'common_name' => 'CommonName',
		];

		foreach($this->country_data as &$country) {
			$local = [];
			foreach($vars as $key => $gokey) {
				if (!isset($country[$key])) continue; // not existing
				$v = $translate($country[$key]);
				if (is_null($v)) continue; // not translated
				$local[$gokey] = $v;
			}
			if (!$local) continue; // no translations for this country

			// output translations
			fwrite($fp, "\t".self::goescape($country['alpha_2']).": Translated{\n");
			foreach($local as $k => $v) {
				fwrite($fp, "\t\t$k: ".self::goescape($v).",\n");
			}
			fwrite($fp, "\t},\n");
		}

		fwrite($fp, "}\n");
		fclose($fp);
	}

	public function genLanguagesIndex() {
		// write language index
		$fp = fopen('countrynames/index-lng.go', 'w');
		fwrite($fp, "package countrynames\n\n");
		fwrite($fp, "var Locale = map[string]map[string]Translated{\n");

		foreach($this->lngIndex as $k => $v)
			fwrite($fp, "\t".self::goescape($k).": $v,\n");

		fwrite($fp, "}\n");
		fclose($fp);

		$fp = fopen('countrynames/index-lngtag.go', 'w');
		fwrite($fp, "package countrynames\n\n");
		fwrite($fp, "import \"golang.org/x/text/language\"\n\n");
		fwrite($fp, "var LocaleByTag = map[language.Tag]map[string]Translated{\n");
		foreach($this->golng as $lng => $tagname) {
			if (!isset($this->lngIndex[$lng])) continue;
			fwrite($fp, "\tlanguage.$tagname: ".$this->lngIndex[$lng].",\n");
		}
		fwrite($fp, "}\n");
		fclose($fp);
	}

	public static function asciify($var) {
		return transliterator_transliterate('Latin-ASCII', transliterator_transliterate('Any-Latin', $var));
	}

	public static function goescape($str) {
		if (is_null($str)) return '""';
		return '"'.addcslashes($str, "\0..\37\"\\").'"';
	}
}
