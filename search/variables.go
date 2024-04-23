package search

import "regexp"

const ep = "https://www.airbnb.com/api/v3/StaysSearch/d4d9503616dc72ab220ed8dcf17f166816dccb2593e7b4625c91c3fce3a3b3d6"

var regexNumber = regexp.MustCompile(`\d+`)
var treament = []string{
	"feed_map_decouple_m11_treatment",
	"stays_search_rehydration_treatment_desktop",
	"stays_search_rehydration_treatment_moweb",
	"selective_query_feed_map_homepage_desktop_treatment",
	"selective_query_feed_map_homepage_moweb_treatment",
}
