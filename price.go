package gobnb

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/johnbalvin/gobnb/trace"
)

func (pdi PriceDependencyInput) GetPrice(currency string, cookies []*http.Cookie, proxyURL *url.URL) (Price, error) {
	urlParsed, err := url.Parse(ep)
	if err != nil {
		return Price{}, trace.NewOrAdd(1, "main", "PriceDependencyInput Price", err, "")
	}
	query := url.Values{}
	query.Add("operationName", "StaysPdpSections")
	query.Add("locale", "en")
	query.Add("currency", currency)
	entension := metadatExtension{
		PersistedQuery: persistedQuery{
			Version:    1,
			Sha256Hash: "e6a7821cf0f78dfc0baab6fd111027eb2976355f2aecbb84bc2086ee6e57161b",
		},
	}
	dataRawExtension, err := json.Marshal(entension)
	if err != nil {
		return Price{}, trace.NewOrAdd(2, "main", "getMetadata", err, "")
	}
	variablesData := metadataVariables{
		ID: pdi.ProducID,
		PdpSectionsRequest: PdpSectionsRequest{
			Adults:                        "1",
			BypassTargetings:              false,
			CategoryTag:                   nil,
			CauseId:                       nil,
			Children:                      nil,
			DisasterId:                    nil,
			DiscountedGuestFeeVersion:     nil,
			DisplayExtensions:             nil,
			FederatedSearchId:             nil,
			ForceBoostPriorityMessageType: nil,
			Infants:                       nil,
			InteractionType:               nil,
			Layouts:                       []string{"SIDEBAR", "SINGLE_COLUMN"},
			Pets:                          0,
			PdpTypeOverride:               nil,
			PhotoId:                       nil,
			Preview:                       false,
			PreviousStateCheckIn:          nil,
			PreviousStateCheckOut:         nil,
			PriceDropSource:               nil,
			PrivateBooking:                false,
			PromotionUuid:                 nil,
			RelaxedAmenityIds:             nil,
			SearchId:                      nil,
			SelectedCancellationPolicyId:  nil,
			SelectedRatePlanId:            nil,
			SplitStays:                    nil,
			StaysBookingMigrationEnabled:  false,
			TranslateUgc:                  nil,
			UseNewSectionWrapperApi:       false,
			SectionIds: []string{
				"CANCELLATION_POLICY_PICKER_MODAL", "BOOK_IT_CALENDAR_SHEET", "POLICIES_DEFAULT", "BOOK_IT_SIDEBAR", "URGENCY_COMMITMENT_SIDEBAR",
				"BOOK_IT_NAV", "BOOK_IT_FLOATING_FOOTER", "EDUCATION_FOOTER_BANNER", "URGENCY_COMMITMENT", "EDUCATION_FOOTER_BANNER_MODAL"},
			CheckIn:        nil,
			CheckOut:       nil,
			P3ImpressionId: pdi.ImpresionID,
		},
	}
	dataRawVariables, err := json.Marshal(variablesData)
	if err != nil {
		return Price{}, trace.NewOrAdd(2, "main", "PriceDependencyInput Price", err, "")
	}
	query.Add("variables", string(dataRawVariables))
	query.Add("extensions", string(dataRawExtension))
	urlParsed.RawQuery = query.Encode()
	req, err := http.NewRequest("GET", urlParsed.String(), nil)
	if err != nil {
		return Price{}, trace.NewOrAdd(1, "main", "PriceDependencyInput Price", err, "")
	}
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("Accept-Language", "en")
	req.Header.Add("Sec-Ch-Ua", `"Not_A Brand";v="8", "Chromium";v="120", "Google Chrome";v="120"`)
	//req.Header.Add("Referer", `https://www.airbnb.com/rooms/793412589557290987`)
	req.Header.Add("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Add("Content-Type", `application/json`)
	req.Header.Add("Sec-Ch-Ua-Platform", `"Windows"`)
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("X-Airbnb-Api-Key", pdi.ApiKey)
	req.Header.Add("Sec-Fetch-User", "?1")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	transport := &http.Transport{
		MaxIdleConnsPerHost: 30,
		DisableKeepAlives:   true,
		TLSClientConfig: &tls.Config{
			Renegotiation:      tls.RenegotiateOnceAsClient,
			InsecureSkipVerify: true,
		},
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
		return Price{}, trace.NewOrAdd(3, "main", "PriceDependencyInput Price", err, "")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Price{}, trace.NewOrAdd(4, "main", "PriceDependencyInput Price", err, "")
	}
	var metadataData2 metadataData2
	if err := json.Unmarshal(body, &metadataData2); err != nil {
		return Price{}, trace.NewOrAdd(5, "main", "PriceDependencyInput Price", err, "")
	}
	if resp.StatusCode != 200 {
		errData := fmt.Sprintf("status: %d headers: %+v", resp.StatusCode, resp.Header)
		return Price{}, trace.NewOrAdd(6, "main", "PriceDependencyInput Price", trace.ErrStatusCode, errData)
	}
	//log.Printf("price 1: %+v\n", metadataData2.Data.Presentation.StayProductDetailPage.Sections.Metadata.BookingPrefetchData.P3_display_rate)
	for _, section := range metadataData2.Data.Presentation.StayProductDetailPage.Sections.Section {
		if section.Section9.Typename == "BookItSection" {
			pr := section.Section9.StructuredDisplayPrice.PrimaryLine
			priceCurrency, priceConverted, err := parsePriceSymbol(pr.Price)
			if err != nil {
				return Price{}, trace.NewOrAdd(5, "main", "PriceDependencyInput Price", err, "")
			}
			return Price{
				Amount:         priceConverted,
				CurrencySymbol: priceCurrency,
				Qualifier:      pr.Qualifier,
			}, nil
		}
	}
	return Price{}, nil
}
