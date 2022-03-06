# countries database

Based on ISO 3166-1 

Translations for country names/etc fetched from https://salsa.debian.org/iso-codes-team/iso-codes

# Usage

You can get data directly for a given country, such as like this:

```go
info := countrydb.France;
return info.ISO3166_Alpha2;
```

## Country lookup by iso code

```go
country, found := countrydb.ByAlpha2[code];
if found {
	...
}
```
