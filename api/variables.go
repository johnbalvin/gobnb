package api

import "regexp"

const ep = "https://www.airbnb.com"

var (
	regxApiKey = regexp.MustCompile(`"key":".+?"`)
)
