package search

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/johnbalvin/gobnb/api"
	"github.com/johnbalvin/gobnb/trace"
	"github.com/johnbalvin/gobnb/utils"
)

func (input InputData) SearchFirstPage(currency string, proxyURL *url.URL) ([]Data, error) {
	apiKey, err := api.Get(proxyURL)
	if err != nil {
		return nil, trace.NewOrAdd(1, "search", "SearchFirstPage", err, "")
	}
	log.Printf("key: *%s*\n", apiKey)
	result, err := input.search("", currency, apiKey, proxyURL)
	if err != nil {
		return nil, trace.NewOrAdd(2, "search", "SearchFirstPage", err, "")
	}
	return result.standardize(), nil
}
func (input InputData) SearchAll(currency string, proxyURL *url.URL) ([]Data, error) {
	apiKey, err := api.Get(proxyURL)
	if err != nil {
		return nil, trace.NewOrAdd(1, "search", "SearchAll", err, "")
	}
	var allResults []Data
	var cursor string
	for {
		resultsRaw, err := input.search(cursor, currency, apiKey, proxyURL)
		if err != nil {
			errData := trace.NewOrAdd(2, "search", "SearchAll", err, "")
			log.Println(errData)
			break
		}
		results := resultsRaw.standardize()
		fmt.Printf("Results len: %d\n", len(results))
		cursor = resultsRaw.PaginationInfo.NextPageCursor
		allResults = append(allResults, results...)
		if len(results) == 0 || cursor == "" {
			break
		}
	}
	return allResults, nil
}
func (input InputData) search(cursor, currency, apiKey string, proxyURL *url.URL) (rootdatapresentationstayssearchResults, error) {
	checkinS := getStringDate(input.Check.In)
	checkoutS := getStringDate(input.Check.Out)
	hours := input.Check.Out.Sub(input.Check.In).Hours()
	days := int(hours / 24)
	urlParsed, err := url.Parse(ep)
	if err != nil {
		return rootdatapresentationstayssearchResults{}, trace.NewOrAdd(1, "search", "search", err, "")
	}
	query := url.Values{}
	query.Add("operationName", "StaysSearch")
	query.Add("locale", "en")
	query.Add("currency", currency)
	urlParsed.RawQuery = query.Encode()
	urlToUse := urlParsed.String()
	rawParames := []rawParam{
		{FilterName: "cdnCacheSafe", FilterValues: []string{"false"}},
		{FilterName: "channel", FilterValues: []string{"EXPLORE"}},
		{FilterName: "checkin", FilterValues: []string{checkinS}},
		{FilterName: "checkout", FilterValues: []string{checkoutS}},
		{FilterName: "datePickerType", FilterValues: []string{"calendar"}},
		{FilterName: "flexibleTripLengths", FilterValues: []string{"one_week"}},
		{FilterName: "itemsPerGrid", FilterValues: []string{"50"}}, //if you read this, this is items returned number, this can bex exploited  ;)
		{FilterName: "monthlyLength", FilterValues: []string{"3"}},
		{FilterName: "monthlyStartDate", FilterValues: []string{"2024-02-01"}},
		{FilterName: "neLat", FilterValues: []string{fmt.Sprintf("%f", input.Coordinates.Ne.Latitude)}},
		{FilterName: "neLng", FilterValues: []string{fmt.Sprintf("%f", input.Coordinates.Ne.Longitud)}},
		{FilterName: "placeId", FilterValues: []string{"ChIJpTeBx6wjq5oROJeXkPCSSSo"}},
		{FilterName: "priceFilterInputType", FilterValues: []string{"0"}},
		{FilterName: "priceFilterNumNights", FilterValues: []string{fmt.Sprintf("%d", days)}},
		{FilterName: "query", FilterValues: []string{"Galapagos Island, Ecuador"}},
		{FilterName: "screenSize", FilterValues: []string{"large"}},
		{FilterName: "refinementPaths", FilterValues: []string{"/homes"}},
		{FilterName: "searchByMap", FilterValues: []string{"true"}},
		{FilterName: "swLat", FilterValues: []string{fmt.Sprintf("%f", input.Coordinates.Sw.Latitude)}},
		{FilterName: "swLng", FilterValues: []string{fmt.Sprintf("%f", input.Coordinates.Sw.Longitud)}},
		{FilterName: "tabId", FilterValues: []string{"home_tab"}},
		{FilterName: "version", FilterValues: []string{"1.8.3"}},
		{FilterName: "zoomLevel", FilterValues: []string{fmt.Sprintf("%d", input.ZoomValue)}},
	}
	inputData := searchRequest{
		OperationName: "StaysSearch",
		Extensions: extensions{
			PersistedQuery: persistedQuery{
				Version:    1,
				Sha256Hash: "d4d9503616dc72ab220ed8dcf17f166816dccb2593e7b4625c91c3fce3a3b3d6",
			},
		},
		Variables: variables{
			IncludeMapResults: true,
			IsLeanTreatment:   false,
			StaysMapSearchRequestV2: staysMapSearchRequestV2{
				Cursor:            cursor,
				RequestedPageType: "STAYS_SEARCH",
				MetadataOnly:      false,
				Source:            "structured_search_input_header",
				SearchType:        "user_map_move",
				TreatmentFlags:    treament,
				RawParams:         rawParames,
			},
			StaysSearchRequest: staysSearchRequest{
				Cursor:            cursor,
				MaxMapItems:       9999,
				RequestedPageType: "STAYS_SEARCH",
				MetadataOnly:      false,
				Source:            "structured_search_input_header",
				SearchType:        "user_map_move",
				TreatmentFlags:    treament,
				RawParams:         rawParames,
			},
		},
	}
	rawData, err := json.Marshal(inputData)
	if err != nil {
		return rootdatapresentationstayssearchResults{}, trace.NewOrAdd(2, "search", "search", err, "")
	}
	req, err := http.NewRequest("POST", urlToUse, bytes.NewReader(rawData))
	if err != nil {
		return rootdatapresentationstayssearchResults{}, trace.NewOrAdd(3, "search", "search", err, "")
	}
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("Accept-Language", "en")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Sec-Ch-Ua", `"Not_A Brand";v="8", "Chromium";v="120", "Google Chrome";v="120"`)
	req.Header.Add("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Add("X-Airbnb-Api-Key", apiKey)
	req.Header.Add("Sec-Ch-Ua-Platform", `"Windows"`)
	req.Header.Add("Sec-Fetch-Dest", "document")
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-Site", "none")
	req.Header.Add("Sec-Fetch-User", "?1")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	transport := &http.Transport{
		MaxIdleConnsPerHost: 30,
		DisableKeepAlives:   true,
	}
	if proxyURL != nil {
		transport.Proxy = http.ProxyURL(proxyURL)
	}
	client := &http.Client{
		Timeout: time.Minute,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: transport,
	}
	resp, err := client.Do(req)
	if err != nil {
		return rootdatapresentationstayssearchResults{}, trace.NewOrAdd(4, "search", "search", err, "")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return rootdatapresentationstayssearchResults{}, trace.NewOrAdd(5, "search", "search", err, "")
	}
	if resp.StatusCode != 200 {
		errData := fmt.Sprintf("status: %d headers: %+v", resp.StatusCode, resp.Header)
		return rootdatapresentationstayssearchResults{}, trace.NewOrAdd(6, "search", "search", trace.ErrStatusCode, errData)
	}
	body = utils.RemoveSpaceByte(body) //some values are returned with weird empty values
	var data root
	if err := json.Unmarshal(body, &data); err != nil {
		return rootdatapresentationstayssearchResults{}, trace.NewOrAdd(7, "search", "search", err, "")
	}
	return data.Data.Presentation.StaysSearch.Results, nil
}
