package countrynames

import "golang.org/x/text/language"

var LocaleByTag = map[language.Tag]map[string]Translated{
	language.Afrikaans:           Afrikaans,
	language.Amharic:             Amharic,
	language.Arabic:              Arabic,
	language.Azerbaijani:         Azerbaijani,
	language.Bulgarian:           Bulgarian,
	language.Bengali:             Bengali,
	language.Catalan:             Catalan,
	language.Czech:               Czech,
	language.Danish:              Danish,
	language.German:              German,
	language.Greek:               Greek,
	language.English:             English,
	language.Spanish:             Spanish,
	language.Estonian:            Estonian,
	language.Persian:             Persian,
	language.Finnish:             Finnish,
	language.Filipino:            Filipino,
	language.French:              French,
	language.Gujarati:            Gujarati,
	language.Hebrew:              Hebrew,
	language.Hindi:               Hindi,
	language.Croatian:            Croatian,
	language.Hungarian:           Hungarian,
	language.Armenian:            Armenian,
	language.Indonesian:          Indonesian,
	language.Icelandic:           Icelandic,
	language.Italian:             Italian,
	language.Japanese:            Japanese,
	language.Georgian:            Georgian,
	language.Kazakh:              Kazakh,
	language.Khmer:               Khmer,
	language.Kannada:             Kannada,
	language.Korean:              Korean,
	language.Kirghiz:             Kirghiz,
	language.Lao:                 Lao,
	language.Lithuanian:          Lithuanian,
	language.Latvian:             Latvian,
	language.Macedonian:          Macedonian,
	language.Malayalam:           Malayalam,
	language.Mongolian:           Mongolian,
	language.Marathi:             Marathi,
	language.Malay:               Malay,
	language.Burmese:             Burmese,
	language.Nepali:              Nepali,
	language.Dutch:               Dutch,
	language.Punjabi:             Punjabi,
	language.Polish:              Polish,
	language.Portuguese:          Portuguese,
	language.BrazilianPortuguese: BrazilianPortuguese,
	language.Romanian:            Romanian,
	language.Russian:             Russian,
	language.Sinhala:             Sinhala,
	language.Slovak:              Slovak,
	language.Slovenian:           Slovenian,
	language.Albanian:            Albanian,
	language.Serbian:             Serbian,
	language.SerbianLatin:        SerbianLatin,
	language.Swedish:             Swedish,
	language.Swahili:             Swahili,
	language.Tamil:               Tamil,
	language.Telugu:              Telugu,
	language.Thai:                Thai,
	language.Turkish:             Turkish,
	language.Ukrainian:           Ukrainian,
	language.Urdu:                Urdu,
	language.Uzbek:               Uzbek,
	language.Vietnamese:          Vietnamese,
	language.SimplifiedChinese:   SimplifiedChinese,
	language.TraditionalChinese:  TraditionalChinese,
	language.Zulu:                Zulu,
}
