package countrydb

type Country struct {
	Name           string
	UniqueName     string
	OfficialName   string // optional
	CommonName     string // optional
	ISO3166_Alpha2 string
	ISO3166_Alpha3 string
	Numeric        int
	CcTLD          string
	FIPS           string // FIPS 10 / GEC code
	Currency       string // 3 letters currency code (may be empty)
	ICANN_Region   string // icann region as defined by https://meetings.icann.org/en/regions
}
