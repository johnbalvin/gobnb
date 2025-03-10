package details

import "regexp"

const ep = "https://www.airbnb.com/api/v3/StaysPdpSections/6f2c582da19b486271d60c4b19e7bdd1147184662f1f4e9a83b08211a73d7343"

var (
	regxApiKey    = regexp.MustCompile(`"key":".+?"`)
	regexLanguage = regexp.MustCompile(`"language":".+?"`)
	regexListing  = regexp.MustCompile(`"id":"\d+?","listingObjType"`)
	regexNumber   = regexp.MustCompile(`\d+`)
)
