package api

import "regexp"

const ep = "https://www.airbnb.com"

var (
	regxApiKey = regexp.MustCompile(`"api_config":{"key":".+?"`)
)
