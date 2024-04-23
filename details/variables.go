package details

import "regexp"

const ep = "https://www.airbnb.com/api/v3/StaysPdpSections"

var (
	regxApiKey    = regexp.MustCompile(`"key":".+?"`)
	regxPrice     = regexp.MustCompile(`\d.+`)
	regexLanguage = regexp.MustCompile(`"language":".+?"`)
	regexListing  = regexp.MustCompile(`"id":"\d+?","listingObjType"`)
	regexNumber   = regexp.MustCompile(`\d+`)
)
