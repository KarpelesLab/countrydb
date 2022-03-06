package countrydb

type Country struct {
	Name           string
	UniqueName     string
	OfficialName   string // optional
	CommonName     string // optional
	ISO3166_Alpha2 string
	ISO3166_Alpha3 string
	Numeric        int
}
