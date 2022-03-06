package countrydb

var Localesi = map[*Country]*Translated{
	Aruba: &Translated{
		Name: "ඇරූබා",
	},
	Afghanistan: &Translated{
		Name:         "ඇෆ්ගනිස්තානය",
		OfficialName: "ඇෆ්ගනිස්තාන ඉස්ලාමීය ජනරජය",
	},
	Angola: &Translated{
		Name:         "ඇන්ගෝලාව",
		OfficialName: "ඇන්ගෝලා ජනරජය",
	},
	Anguilla: &Translated{
		Name: "ඇන්ගුයිලා",
	},
	AlandIslands: &Translated{
		Name: "ඒලන්ඩ් දූපත්",
	},
	Albania: &Translated{
		Name:         "ඇල්බේනියාව",
		OfficialName: "ඇල්බේනියා ජනරජය",
	},
	Andorra: &Translated{
		Name:         "ඇන්ඩෝරා",
		OfficialName: "ඇන්ඩෝරා පරිපාලනය",
	},
	UnitedArabEmirates: &Translated{
		Name: "එක්සත් අරාබි එමීර්",
	},
	Argentina: &Translated{
		Name:         "ආජන්ටිනාව",
		OfficialName: "ආජන්ටිනා ජනරජය",
	},
	Armenia: &Translated{
		Name:         "ආමේනියා",
		OfficialName: "ආමේනියා ජනරජය",
	},
	AmericanSamoa: &Translated{
		Name: "ඇමරිකානු සැමෝවා",
	},
	Antarctica: &Translated{
		Name: "ඇන්ටාක්ටිකා",
	},
	FrenchSouthernTerritories: &Translated{
		Name: "ප්‍රංශ දකුණු ප්‍රදේශ",
	},
	AntiguaandBarbuda: &Translated{
		Name: "ඇන්ටිගුවා හා බාබුඩා",
	},
	Australia: &Translated{
		Name: "ඔස්ට්‍රේලියාව",
	},
	Austria: &Translated{
		Name:         "ඔස්ට්‍රියාව",
		OfficialName: "ඔස්ට්‍රියා ජනරජය",
	},
	Azerbaijan: &Translated{
		Name:         "අසර්බයිජාන්",
		OfficialName: "අසර්බයිජාන් ජනරජය",
	},
	Burundi: &Translated{
		Name:         "බුරුන්ඩි",
		OfficialName: "බුරුන්ඩි ජනරජය",
	},
	Belgium: &Translated{
		Name:         "බෙල්ජියම",
		OfficialName: "බෙල්ජියම් රාජධානිය",
	},
	Benin: &Translated{
		Name:         "බෙනින්",
		OfficialName: "බෙනින් ජනරජය",
	},
	Bonaire: &Translated{
		Name:         "බොනාරේ. සිනෙට් එස්ටාටියස් හා සාබා",
		OfficialName: "බොනාරේ. සිනෙට් එස්ටාටියස් හා සාබා",
	},
	BurkinaFaso: &Translated{
		Name: "බර්කිනා ෆාසෝ",
	},
	Bangladesh: &Translated{
		Name:         "බංගලාදේශය",
		OfficialName: "බංගලාදේශ මහජන ජනරජය",
	},
	Bulgaria: &Translated{
		Name:         "බල්ගේරියා",
		OfficialName: "බල්ගේරියා ජනරජය",
	},
	Bahrain: &Translated{
		Name:         "බහරේන්",
		OfficialName: "බහරේන් රාජධානිය",
	},
	Bahamas: &Translated{
		Name:         "බහමාස්",
		OfficialName: "බහමාස් පොදුරාජ්‍යය",
	},
	BosniaandHerzegovina: &Translated{
		Name:         "බොස්නියාව හා හර්සගෝවිනියාව",
		OfficialName: "බොස්නියා හා හෙර්ස්ගොවිනා ජනරජය",
	},
	SaintBarthelemy: &Translated{
		Name: "ශාන්ත බර්තොලමියු",
	},
	Belarus: &Translated{
		Name:         "බෙලාරුස්",
		OfficialName: "බෙලාරුස් ජනරජය",
	},
	Belize: &Translated{
		Name: "බිලිසේ",
	},
	Bermuda: &Translated{
		Name: "බර්මියුඩා",
	},
	Bolivia: &Translated{
		Name:         "බොලීවියා බහුජාතීය රාජ්‍යය",
		OfficialName: "බොලීවියා බහුජාතීය රාජ්‍යය",
		CommonName:   "බොලීවියාව",
	},
	Brazil: &Translated{
		Name:         "බ්‍රසිලය",
		OfficialName: "බ්‍රසීල ෆෙඩරල් ජනරජය",
	},
	Barbados: &Translated{
		Name: "බාබඩෝස්",
	},
	BruneiDarussalam: &Translated{
		Name: "බෲනායි දරුසලම්",
	},
	Bhutan: &Translated{
		Name:         "භූතානය",
		OfficialName: "භූතාන රාජධානිය",
	},
	BouvetIsland: &Translated{
		Name: "බෝවෙට් දූපත",
	},
	Botswana: &Translated{
		Name:         "බොට්ස්වානා",
		OfficialName: "බොස්ට්වානා ජනරජය",
	},
	CentralAfricanRepublic: &Translated{
		Name: "මධ්‍යම අප්‍රිකානු ජනරජය",
	},
	Canada: &Translated{
		Name: "කැනඩාව",
	},
	CocosKeelingIslands: &Translated{
		Name: "කොකෝවා (කීලිං) දූපත්",
	},
	Switzerland: &Translated{
		Name:         "ස්විට්සර්ලන්තය",
		OfficialName: "ස්විස් සංගමය",
	},
	Chile: &Translated{
		Name:         "චිලී",
		OfficialName: "චිලී ජනරජය",
	},
	China: &Translated{
		Name:         "චීනය",
		OfficialName: "චීන මහජන ජනරජය",
	},
	CotedIvoire: &Translated{
		Name:         "කෝටේ ඩ්ලවොයිරා",
		OfficialName: "කෝටේ ඩ්ලවොයිරා ජනරජය",
	},
	Cameroon: &Translated{
		Name:         "කැමරූන්",
		OfficialName: "කැමරූන් ජනරජය",
	},
	DemocraticCongo: &Translated{
		Name: "කොන්ගෝ ප්‍රජාතන්ත්‍රවාදී ජනරජය",
	},
	Congo: &Translated{
		Name:         "කොන්ගෝව",
		OfficialName: "කොන්ගෝ ජනරජය",
	},
	CookIslands: &Translated{
		Name: "කූක් දූපත්",
	},
	Colombia: &Translated{
		Name:         "කොලොම්බියාව",
		OfficialName: "කොලොම්බියා ජනරජය",
	},
	Comoros: &Translated{
		Name:         "කොමොරොස්",
		OfficialName: "කොමොරොස් සංගමය",
	},
	CaboVerde: &Translated{
		Name:         "කේප් වර්ඩෙe",
		OfficialName: "කේප් වර්ඩේ ජනරජය",
	},
	CostaRica: &Translated{
		Name:         "කොස්ටා රීකා",
		OfficialName: "කොස්ටා රිකා ජනරජය",
	},
	Cuba: &Translated{
		Name:         "කියුබා",
		OfficialName: "කියුබා ජනරජය",
	},
	Curacao: &Translated{
		Name:         "කුරාකෝ",
		OfficialName: "කුරාකෝ",
	},
	ChristmasIsland: &Translated{
		Name: "ක්‍රිස්මස් දූපත්",
	},
	CaymanIslands: &Translated{
		Name: "සයිමන් දූපත්",
	},
	Cyprus: &Translated{
		Name:         "සයිප්‍රස්",
		OfficialName: "සයිප්‍රස් ජනරජය",
	},
	Czechia: &Translated{
		OfficialName: "චෙච් ජනරජය",
	},
	Germany: &Translated{
		Name:         "ජර්මනිය",
		OfficialName: "ජර්මානු ෆෙඩරල් ජනරජය",
	},
	Djibouti: &Translated{
		Name:         "ජිල්බූට්",
		OfficialName: "ජිල්බූටි ජනරජය",
	},
	Dominica: &Translated{
		Name:         "ඩොමිනිකාව",
		OfficialName: "ඩොමිනිකා පොදුරාජ්‍යය",
	},
	Denmark: &Translated{
		Name:         "ඩෙන්මාකය",
		OfficialName: "ඩෙන්මාක රාජධානිය",
	},
	DominicanRepublic: &Translated{
		Name: "ඩොමිනිකානු ජනරජය",
	},
	Algeria: &Translated{
		Name:         "ඇල්ජීරියා",
		OfficialName: "මහජන ප්‍රජාතාන්ත්‍රික ඇල්බේනියා ජනරජය",
	},
	Ecuador: &Translated{
		Name:         "ඉක්වදෝරය",
		OfficialName: "ඉක්වදෝර් ජනරජය",
	},
	Egypt: &Translated{
		Name:         "ඊජිප්තුව",
		OfficialName: "ඊජිප්තු අරාබි ජනරජය",
	},
	Eritrea: &Translated{
		Name:         "එරිත්‍රියා",
		OfficialName: "ඇමරිකා එක්සත් ජනපද",
	},
	WesternSahara: &Translated{
		Name: "බටහිර සහරාව",
	},
	Spain: &Translated{
		Name:         "ස්පාඤ්ඤය",
		OfficialName: "ස්පාඤ්ඤ රාජධානිය",
	},
	Estonia: &Translated{
		Name:         "එස්තෝනියා",
		OfficialName: "එස්තෝනියා ජනරජය",
	},
	Ethiopia: &Translated{
		Name:         "ඉතියෝපියාව",
		OfficialName: "ෆෙඩරල් ප්‍රජාතන්ත්‍රවාදී ඉතියෝපියානු ජනරජය",
	},
	Finland: &Translated{
		Name:         "ෆින්ලන්තය",
		OfficialName: "ෆින්ලන්ත ජනරජය",
	},
	Fiji: &Translated{
		Name:         "ෆීජි",
		OfficialName: "ෆීජි ජනරජය",
	},
	FalklandIslandsMalvinas: &Translated{
		Name: "ෆෝක්ලන්ඩ් දූපත් (මැල්විනා)",
	},
	France: &Translated{
		Name:         "ප්‍රංශය",
		OfficialName: "ප්‍රංශ ජනරජය",
	},
	FaroeIslands: &Translated{
		Name: "ෆාරෝ දූපත්",
	},
	Micronesia: &Translated{
		Name:         "මයික්‍රෝසියා ෆෙඩරල් ජනපද",
		OfficialName: "මයික්‍රෝසියා ෆෙඩරල් ජනපද",
	},
	Gabon: &Translated{
		Name:         "ගාබෝන්",
		OfficialName: "ගබෝනීස් ජනරජය",
	},
	UnitedKingdom: &Translated{
		Name:         "එක්සත් රාජධානිය",
		OfficialName: "මහා බ්‍රිතාන්‍ය හා උතුරු අයර්ලන්තයේ එක්සත් රාජධානිය",
	},
	Georgia: &Translated{
		Name: "ජෝජියාව",
	},
	Guernsey: &Translated{
		Name: "ගුවන්සි",
	},
	Ghana: &Translated{
		Name:         "ඝානාව",
		OfficialName: "ඝානා ජනරජය",
	},
	Gibraltar: &Translated{
		Name: "ජිබෝල්ටා",
	},
	Guinea: &Translated{
		Name:         "ගිනියාව",
		OfficialName: "ගිනියා ජනරජය",
	},
	Guadeloupe: &Translated{
		Name: "ගුවාඩිලෝප්",
	},
	Gambia: &Translated{
		Name:         "ගැම්බියාව",
		OfficialName: "සැම්බියා ජනරජය",
	},
	GuineaBissau: &Translated{
		Name:         "ගිනියා-බිසව්",
		OfficialName: "ගිනියා-බිසව් ජනරජය",
	},
	EquatorialGuinea: &Translated{
		Name:         "ගිනියා සමානාත්මය",
		OfficialName: "ගිනියා සමානාත්ම ජනරජය",
	},
	Greece: &Translated{
		Name:         "ග්‍රීසිය",
		OfficialName: "හෙලනික් ජනරජය",
	},
	Grenada: &Translated{
		Name: "ග්‍රෙනාඩා",
	},
	Greenland: &Translated{
		Name: "ග්‍රීන්ලන්තය",
	},
	Guatemala: &Translated{
		Name:         "ගෝතමාලාව",
		OfficialName: "ගෝතමාලා ජනරජය",
	},
	FrenchGuiana: &Translated{
		Name: "ප්‍රංශ ගිනියාව",
	},
	Guam: &Translated{
		Name: "ගුවාම්",
	},
	Guyana: &Translated{
		Name:         "ගයනා",
		OfficialName: "ගයනා ජනරජය",
	},
	HongKong: &Translated{
		Name:         "හොංග් කොංග්",
		OfficialName: "චීන විශේෂ පාලිත හොංග් කොංග් ප්‍රදේශය",
	},
	HeardIslandandMcDonaldIslands: &Translated{
		Name: "හර්ඩ් හා මැක්ඩොනල්ඩ් දූපත්",
	},
	Honduras: &Translated{
		Name:         "හොන්ඩුරාස්",
		OfficialName: "හොන්ඩුරාස් ජනරජය",
	},
	Croatia: &Translated{
		Name:         "ක්‍රෝශියා",
		OfficialName: "කෝශ්‍රියා ජනරජය",
	},
	Haiti: &Translated{
		Name:         "හයිටි",
		OfficialName: "හයිටි ජනරජය",
	},
	Hungary: &Translated{
		Name:         "හංගේරියාව",
		OfficialName: "හංගේරියාව",
	},
	Indonesia: &Translated{
		Name:         "ඉන්දුනීසියාව",
		OfficialName: "ඉන්දුනීසියා ජනරජය",
	},
	IsleofMan: &Translated{
		Name: "මාන් දිවයින",
	},
	India: &Translated{
		Name:         "ඉන්දියාව",
		OfficialName: "ඉන්දියා ජනරජය",
	},
	BritishIndianOceanTerritory: &Translated{
		Name: "බ්‍රිතාන්‍ය ඉන්දීය මුහුදු ප්‍රදේශ",
	},
	Ireland: &Translated{
		Name: "අයර්ලන්තය",
	},
	Iran: &Translated{
		Name:         "ඉරාන ඉස්ලාමීය ජනරජය",
		OfficialName: "ඉරාන ඉස්ලාමීය ජනරජය",
	},
	Iraq: &Translated{
		Name:         "ඉරාකය",
		OfficialName: "ඉරාක ජනරජය",
	},
	Iceland: &Translated{
		Name:         "අයිස්ලන්තය",
		OfficialName: "අයිස්ලන්ත ජනරජය",
	},
	Israel: &Translated{
		Name:         "ඊශ්‍රායලය",
		OfficialName: "ඊශ්‍රායල ජනපදය",
	},
	Italy: &Translated{
		Name:         "ඉතාලිය",
		OfficialName: "ඉතාලි ජනරජය",
	},
	Jamaica: &Translated{
		Name: "ජැමෙයිකාව",
	},
	Jersey: &Translated{
		Name: "ජර්සි",
	},
	Jordan: &Translated{
		Name:         "ජෝර්දානය",
		OfficialName: "ජෝර්දාන හැශර්මිට් රාජධානිය",
	},
	Japan: &Translated{
		Name: "ජපානය",
	},
	Kazakhstan: &Translated{
		Name:         "කසකස්තානය",
		OfficialName: "කසකස්තාන් ජනරජය",
	},
	Kenya: &Translated{
		Name:         "කෙන්යාව",
		OfficialName: "කෙන්යා ජනරජය",
	},
	Kyrgyzstan: &Translated{
		Name:         "කිර්ගිස්තානය",
		OfficialName: "කිර්ගිස් ජනරජය",
	},
	Cambodia: &Translated{
		Name:         "කාම්බෝජියාව",
		OfficialName: "කාම්බෝජියා රාජධානිය",
	},
	Kiribati: &Translated{
		Name:         "කිරිබටි",
		OfficialName: "කිරිබටි ජනරජය",
	},
	SaintKittsandNevis: &Translated{
		Name: "ශාන්ත කිට්ස් හා නෙවිස්",
	},
	SouthKorea: &Translated{
		Name:       "කොරියා ජනරජය",
		CommonName: "දකුණු අප්‍රිකාව",
	},
	Kuwait: &Translated{
		Name:         "කුවේට්",
		OfficialName: "කුවේට් ජනපදය",
	},
	LaoPeoplesDemocraticRepublic: &Translated{
		Name: "මහජන ප්‍රජාතාන්ත්‍රික ලාවෝ ජනරජය",
	},
	Lebanon: &Translated{
		Name:         "ලෙබනන",
		OfficialName: "ලෙබනන ජනරජය",
	},
	Liberia: &Translated{
		Name:         "ලියිබීරියාව",
		OfficialName: "ලයිබීරියා ජනරජය",
	},
	Libya: &Translated{
		Name:         "ලිබියාව",
		OfficialName: "ලිබියාව",
	},
	SaintLucia: &Translated{
		Name: "ශාන්ත ලුසියා",
	},
	Liechtenstein: &Translated{
		Name:         "ලිච්ටෙන්ස්ටයින්",
		OfficialName: "ලිච්ටෙන්ස්ටයින් පාලනය",
	},
	SriLanka: &Translated{
		Name:         "ශ්‍රී ලංකාව",
		OfficialName: "ශ්‍රීලංකා ප්‍රජාතාන්ත්‍රික සමාජවාදී ජනරජය",
	},
	Lesotho: &Translated{
		Name:         "ලෙස්තෝ",
		OfficialName: "ලෙස්තෝනියා රාජධානිය",
	},
	Lithuania: &Translated{
		Name:         "ලිතුවේනියාව",
		OfficialName: "ලිතුවේනියා ජනරජය",
	},
	Luxembourg: &Translated{
		Name:         "ලක්සම්බර්ග්",
		OfficialName: "මහා ලක්සම්බර්ග්",
	},
	Latvia: &Translated{
		Name:         "ලැත්වියාව",
		OfficialName: "ලැත්වියා ජනරජය",
	},
	Macao: &Translated{
		Name:         "මැකාවෝ",
		OfficialName: "චීන විශේෂ පාලිත මැකාවෝ ප්‍රදේශය",
	},
	SaintMartinFrenchpart: &Translated{
		Name: "ශාන්ත මාටින් (ප්‍රංශ කොටස)",
	},
	Morocco: &Translated{
		Name:         "මොරොක්කෝව",
		OfficialName: "මොරොක්කෝ රාජධානිය",
	},
	Monaco: &Translated{
		Name:         "මොනාකෝ",
		OfficialName: "මොනාකෝ පරිපාලනය",
	},
	Moldova: &Translated{
		Name:         "මෝල්ඩ්වා ජනරජය",
		OfficialName: "මෝල්ඩ්වා ජනරජය",
		CommonName:   "මොල්ඩෝවා",
	},
	Madagascar: &Translated{
		Name:         "මැඩගස්කරය",
		OfficialName: "මැඩගස්කර ජනරජය",
	},
	Maldives: &Translated{
		Name:         "මාල දිවයින",
		OfficialName: "මාල දිවයින් ජනරජය",
	},
	Mexico: &Translated{
		Name:         "මෙක්සිකෝව",
		OfficialName: "එක්සත් මෙක්සිකානු ජනපද",
	},
	MarshallIslands: &Translated{
		Name:         "මාශල් දූපත්",
		OfficialName: "මාශල් දූපත් ජනරජය",
	},
	NorthMacedonia: &Translated{
		Name:         "නව කැලිඩෝනියාව",
		OfficialName: "ලිතුවේනියා ජනරජය",
	},
	Mali: &Translated{
		Name:         "මාලි",
		OfficialName: "මාලි ජනරජය",
	},
	Malta: &Translated{
		Name:         "මෝල්ටා",
		OfficialName: "මෝල්ටා ජනරජය",
	},
	Myanmar: &Translated{
		Name:         "මියන්මාරය",
		OfficialName: "මියන්මාර ජනරජය",
	},
	Montenegro: &Translated{
		Name:         "මොන්ටෙන්ග්‍රෝ",
		OfficialName: "මොන්ටෙන්ග්‍රෝ",
	},
	Mongolia: &Translated{
		Name: "මොංගෝලියාව",
	},
	NorthernMarianaIslands: &Translated{
		Name:         "උතුරු මැරීනා දූපත්",
		OfficialName: "උතුරු මැරීනා පොදුරාජ්‍යය දූපත්",
	},
	Mozambique: &Translated{
		Name:         "මොසැම්බික්",
		OfficialName: "මොසැම්බික් ජනරජය",
	},
	Mauritania: &Translated{
		Name:         "මෝරිටානියා",
		OfficialName: "මඅවුරිතානියා ඉස්ලාමීය ජනරජය",
	},
	Montserrat: &Translated{
		Name: "මොන්ට්සෙරා",
	},
	Martinique: &Translated{
		Name: "මාටිනිකි",
	},
	Mauritius: &Translated{
		Name:         "මවුරිටියස්",
		OfficialName: "මාරිතීස් ජනරජය",
	},
	Malawi: &Translated{
		Name:         "මලාවි",
		OfficialName: "මලාවි ජනරජය",
	},
	Malaysia: &Translated{
		Name: "මලයාසියා",
	},
	Mayotte: &Translated{
		Name: "මයොට්ටේ",
	},
	Namibia: &Translated{
		Name:         "නැම්බියාව",
		OfficialName: "නැම්බියා ජනරජය",
	},
	NewCaledonia: &Translated{
		Name: "නව කැලිඩෝනියාව",
	},
	Niger: &Translated{
		Name:         "නිගර්",
		OfficialName: "නයිගර් ජනරජය",
	},
	NorfolkIsland: &Translated{
		Name: "නෝෆෝක් දූපත",
	},
	Nigeria: &Translated{
		Name:         "නයිජීරියා",
		OfficialName: "නයිජීරියා ෆෙඩරල් ජනරජය",
	},
	Nicaragua: &Translated{
		Name:         "නිකරගුවා",
		OfficialName: "නිකරගුවා ජනරජය",
	},
	Niue: &Translated{
		Name:         "නීයු",
		OfficialName: "නීයු",
	},
	Netherlands: &Translated{
		Name:         "නෙදර්ලන්තය",
		OfficialName: "නෙදර්ලන්ත රාජධානිය",
	},
	Norway: &Translated{
		Name:         "නෝර්වේ",
		OfficialName: "නෝර්වේ රාජධානිය",
	},
	Nepal: &Translated{
		Name:         "නේපාලය",
		OfficialName: "ෆෙඩරල් ප්‍රජාතන්ත්‍රවාදී නේපාල ජනරජය",
	},
	Nauru: &Translated{
		Name:         "නාවුරු",
		OfficialName: "නාඋරු ජනරජය",
	},
	NewZealand: &Translated{
		Name: "නවසීලන්තය",
	},
	Oman: &Translated{
		Name:         "ඕමානය",
		OfficialName: "ඕමාන සුල්තානය",
	},
	Pakistan: &Translated{
		Name:         "පකිස්තානය",
		OfficialName: "පාකිස්තාන ඉස්ලාමීය ජනරජය",
	},
	Panama: &Translated{
		Name:         "පැනමාව",
		OfficialName: "පැනමා ජනරජය",
	},
	Pitcairn: &Translated{
		Name: "පිට්කායිරනව",
	},
	Peru: &Translated{
		Name:         "පේරු",
		OfficialName: "පේරු ජනරජය",
	},
	Philippines: &Translated{
		Name:         "පිලිපීන",
		OfficialName: "පිලිපීන ජනරජය",
	},
	Palau: &Translated{
		Name:         "පාලාව්",
		OfficialName: "පලාව් ජනරජය",
	},
	PapuaNewGuinea: &Translated{
		Name:         "පැපුවා නිව්ගිනියාව",
		OfficialName: "සැමෝවා නිදහස් ජනපදය",
	},
	Poland: &Translated{
		Name:         "පෝලන්තය",
		OfficialName: "පෝලන්ත ජනරජය",
	},
	PuertoRico: &Translated{
		Name: "පුවටෝ රිකෝව",
	},
	NorthKorea: &Translated{
		Name:         "සමාජවාදී මහජන කොරියා ජනරජය",
		OfficialName: "සමාජවාදී මහජන කොරියා ජනරජය",
		CommonName:   "නව කැලිඩෝනියාව",
	},
	Portugal: &Translated{
		Name:         "පෘතුගාලය",
		OfficialName: "පෘතුගීසි ජනරජය",
	},
	Paraguay: &Translated{
		Name:         "පැරගුවේ",
		OfficialName: "පැරගුවේ ජනරජය",
	},
	Palestine: &Translated{
		OfficialName: "ඇමරිකා එක්සත් ජනපද",
	},
	FrenchPolynesia: &Translated{
		Name: "ප්‍රංශ පොලිනීසියාව",
	},
	Qatar: &Translated{
		Name:         "කටාර්",
		OfficialName: "කටාර් ජනපදය",
	},
	Reunion: &Translated{
		Name: "ප්‍රතිසංවිධානය",
	},
	Romania: &Translated{
		Name: "රොමේනියාව",
	},
	RussianFederation: &Translated{
		Name: "රුසියානු සංගමය",
	},
	Rwanda: &Translated{
		Name:         "රුවන්ඩා",
		OfficialName: "රුවන්ඩා ජනරජය",
	},
	SaudiArabia: &Translated{
		Name:         "සවුදි අරාබිය",
		OfficialName: "සවුදි අරාබි රාජධානිය",
	},
	Sudan: &Translated{
		Name:         "සුඩානය",
		OfficialName: "සුඩාන ජනරජය",
	},
	Senegal: &Translated{
		Name:         "සෙනගාලය",
		OfficialName: "සෙනගාල ජනරජය",
	},
	Singapore: &Translated{
		Name:         "සිංගප්පූරුව",
		OfficialName: "සිංගප්පූරු ජනරජය",
	},
	SouthGeorgiaandtheSouthSandwichIslands: &Translated{
		Name: "දකුණු ජෝජියආ හා දකුණු සැන්ඩ්විච් දූපත්",
	},
	SaintHelena: &Translated{
		Name: "ශාන්ත හෙලේනා, ඇසෙන්ශන් හා ට්‍රිස්ටන් ද කුන්හා",
	},
	SvalbardandJanMayen: &Translated{
		Name: "ස්වල්බාඩ් හා ජැන් මෙයන්",
	},
	SolomonIslands: &Translated{
		Name: "සොලමන් දූපත්",
	},
	SierraLeone: &Translated{
		Name:         "සියෙරා ලියොන්",
		OfficialName: "සියෙරා ලියෝන් ජනරජය",
	},
	ElSalvador: &Translated{
		Name:         "එල් සැල්වදෝරය",
		OfficialName: "එල් සැල්වදෝර් ජනරජය",
	},
	SanMarino: &Translated{
		Name:         "ශාන්ත මැරිනෝ",
		OfficialName: "සැන් මැරීනෝ ජනරජය",
	},
	Somalia: &Translated{
		Name:         "සෝමාලියාව",
		OfficialName: "ජර්මානු ෆෙඩරල් ජනරජය",
	},
	SaintPierreandMiquelon: &Translated{
		Name: "ශාන්ත පියරේ හා මිකුලෝන්",
	},
	Serbia: &Translated{
		Name:         "සර්බියාව",
		OfficialName: "සර්බියා ජනරජය",
	},
	SouthSudan: &Translated{
		Name:         "දකුණු සුඩානය",
		OfficialName: "දකුණු සුඩාන් ජනරජය",
	},
	SaoTomeandPrincipe: &Translated{
		Name:         "සාවෝ තෝමේ හා ප්‍රින්සිපි",
		OfficialName: "සාඕ ටෝම් හා ප්‍රින්සිපි ප්‍රජාතාන්ත්‍රික ජනරජය",
	},
	Suriname: &Translated{
		Name:         "සුරිනාම්",
		OfficialName: "සූරිනාම් ජනරජය",
	},
	Slovakia: &Translated{
		Name:         "ස්ලෝවැකියාව",
		OfficialName: "ස්ලෝවැක් ජනරජය",
	},
	Slovenia: &Translated{
		Name:         "ස්ලෝවීනියාව",
		OfficialName: "ස්ලෝවීනියා ජනරජය",
	},
	Sweden: &Translated{
		Name:         "ස්වීඩනය",
		OfficialName: "ස්වීඩන රාජධානිය",
	},
	Eswatini: &Translated{
		OfficialName: "ස්පාඤ්ඤ රාජධානිය",
	},
	SintMaartenDutchpart: &Translated{
		Name:         "සින්ට් මාටෙන් (ඕලන්ද කොටස)",
		OfficialName: "සින්ට් මාටෙන් (ඕලන්ද කොටස)",
	},
	Seychelles: &Translated{
		Name:         "සීශෙල්ස්",
		OfficialName: "සීශෙල්ස් ජනරජය",
	},
	SyrianArabRepublic: &Translated{
		Name: "සිරියානු අරාබි ජනරජය",
	},
	TurksandCaicosIslands: &Translated{
		Name: "තුර්කි හා කායිකෝස් දූපත්",
	},
	Chad: &Translated{
		Name:         "cqDw",
		OfficialName: "චැඩ් ජනරජය",
	},
	Togo: &Translated{
		Name:         "ටෝගෝ",
		OfficialName: "ටෝගෝලීස් ජනරජය",
	},
	Thailand: &Translated{
		Name:         "තායිලන්තය",
		OfficialName: "තායිලන්ත රාජධානිය",
	},
	Tajikistan: &Translated{
		Name:         "ටජිකිස්තානය",
		OfficialName: "ටජිකිස්තාන් ජනරජය",
	},
	Tokelau: &Translated{
		Name: "ටොකෙලාවෝ",
	},
	Turkmenistan: &Translated{
		Name: "තුර්කිමෙනිස්තානය",
	},
	TimorLeste: &Translated{
		Name:         "ටිමෝර්-ලෙස්ටේ",
		OfficialName: "ටිමෝර්-ලෙස්ටේ ප්‍රජාතාන්ත්‍රික ජනරජය",
	},
	Tonga: &Translated{
		Name:         "ටොන්ගා",
		OfficialName: "ටොන්ගෝ රාජධානිය",
	},
	TrinidadandTobago: &Translated{
		Name:         "ට්‍රිනීඩෑඩව හා ටොබාගෝ",
		OfficialName: "ට්‍රිනිඩෑඩ් සහ ටෝබාගෝ ජනරජය",
	},
	Tunisia: &Translated{
		Name:         "ටියුනීසියා",
		OfficialName: "ටියුනීසියා ජනරජය",
	},
	Turkey: &Translated{
		Name:         "තුර්කිය",
		OfficialName: "තුර්කි ජනරජය",
	},
	Tuvalu: &Translated{
		Name: "ටුවාලු",
	},
	Taiwan: &Translated{
		Name:         "චීනයේ තායිවාන ප්‍රදේශය",
		OfficialName: "චීනයේ තායිවාන ප්‍රදේශය",
		CommonName:   "තායිවානය",
	},
	Tanzania: &Translated{
		Name:         "ටැන්සානියා එක්සත් ජනරජය",
		OfficialName: "ටැන්සානියා එක්සත් ජනරජය",
	},
	Uganda: &Translated{
		Name:         "උගන්ඩාව",
		OfficialName: "උගන්ඩා ජනරජය",
	},
	Ukraine: &Translated{
		Name: "යුක්රේනය",
	},
	UnitedStatesMinorOutlyingIslands: &Translated{
		Name: "කුඩා දූපත් එක්සත් ජනපද",
	},
	Uruguay: &Translated{
		Name:         "උරුගුවේ",
		OfficialName: "උරුගුවේ නැගෙනහිර ජනරජය",
	},
	UnitedStates: &Translated{
		Name:         "එක්සත් ජනපද",
		OfficialName: "ඇමරිකා එක්සත් ජනපද",
	},
	Uzbekistan: &Translated{
		Name:         "උස්බෙකිස්තානය",
		OfficialName: "උස්බෙකිස්තාන් ජනරජය",
	},
	HolySeeVaticanCityState: &Translated{
		Name: "ශුද්ධ දැක්ම (වතිකානු නගර රාජ්‍යය)",
	},
	SaintVincentandtheGrenadines: &Translated{
		Name: "ශාන්ත වින්සෙන්ට් හා ග්‍රෙන්ඩිනේස්",
	},
	Venezuela: &Translated{
		Name:         "බොලිවාරියන් වෙනිසියුලා ජනරජය",
		OfficialName: "බොලිවාරියන් වෙනිසියුලා ජනරජය",
		CommonName:   "වෙනිසියුලාව",
	},
	BritishVirginIslands: &Translated{
		Name:         "බ්‍රිතාන්‍ය වර්ජින් දූපත්",
		OfficialName: "බ්‍රිතාන්‍ය වර්ජින් දූපත්",
	},
	VirginIslands: &Translated{
		Name:         "එ.ජ වර්ජින් දූපත්",
		OfficialName: "එක්සත් ජනපද වර්ජින් දූපත්",
	},
	Vietnam: &Translated{
		Name:         "වියෙට්නාමය",
		OfficialName: "වියෙට්නාම සමාජවාදී ජනරජය",
		CommonName:   "වියෙට්නාමය",
	},
	Vanuatu: &Translated{
		Name:         "වනවාටු",
		OfficialName: "වනවාටු ජනරජය",
	},
	WallisandFutuna: &Translated{
		Name: "වලීස් හා ෆුටුනා",
	},
	Samoa: &Translated{
		Name:         "සැමෝවා",
		OfficialName: "සැමෝවා නිදහස් ජනපදය",
	},
	Yemen: &Translated{
		Name:         "යේමනය",
		OfficialName: "යේමන ජනරජය",
	},
	SouthAfrica: &Translated{
		Name:         "දකුණු අප්‍රිකාව",
		OfficialName: "දකුණු අප්‍රිකා ජනරජය",
	},
	Zambia: &Translated{
		Name:         "සැම්බියාව",
		OfficialName: "සැම්බියා ජනරජය",
	},
	Zimbabwe: &Translated{
		Name:         "සිම්බාබ්වේ",
		OfficialName: "සිම්බාබ්වේ ජනරජය",
	},
}
