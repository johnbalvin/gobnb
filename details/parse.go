package details

import (
	"bytes"
	"encoding/json"
	"html"
	"strings"

	"github.com/johnbalvin/gobnb/trace"
	"github.com/johnbalvin/gobnb/utils"

	"github.com/PuerkitoBio/goquery"
)

func ParseBodyDetails(body []byte) (Data, PriceDependencyInput, error) {
	dataRaw, language, apiKey, err := parseBodyDetails(body)
	if err != nil {
		return Data{}, PriceDependencyInput{}, trace.NewOrAdd(1, "main", "ParseBodyDetails", err, "")
	}
	dataFormated := dataRaw.standardize()
	dataFormated.Language = language
	priceDependencyInput := PriceDependencyInput{
		ProducID:    dataRaw.Variables.ID,
		ImpresionID: dataRaw.Variables.PdpSectionsRequest.P3ImpressionId,
		ApiKey:      apiKey,
	}
	return dataFormated, priceDependencyInput, nil
}
func parseBodyDetails(body []byte) (metadataData, string, string, error) {
	reader := bytes.NewReader(body)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return metadataData{}, "", "", trace.NewOrAdd(1, "main", "parseBodyDetails", err, "")
	}
	htmlData, err := doc.Find("#data-deferred-state").Html()
	if err != nil {
		return metadataData{}, "", "", trace.NewOrAdd(2, "main", "parseBodyDetails", err, "")
	}
	htmlData = utils.RemoveSpace(html.UnescapeString(htmlData))
	language := regexLanguage.FindString(htmlData)
	language = strings.ReplaceAll(language, `"language":"`, `"`)
	apiKey := regxApiKey.FindString(string(body))
	apiKey = strings.ReplaceAll(apiKey, `"key":"`, "")
	apiKey = strings.ReplaceAll(apiKey, `"`, "")
	var data niobeMinimalClientDataWrapper
	if err := json.Unmarshal([]byte(htmlData), &data); err != nil {
		return metadataData{}, "", "", trace.NewOrAdd(3, "main", "parseBodyDetails", err, "")
	}
	var datailsData metadataData
	if err := json.Unmarshal(data.NiobeMinimalClientData[0][1], &datailsData); err != nil {
		return metadataData{}, "", "", trace.NewOrAdd(4, "main", "parseBodyDetails", err, "")
	}
	return datailsData, language, apiKey, nil
}
