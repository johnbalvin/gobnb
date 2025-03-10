package details

import (
	"fmt"
	"net/url"

	"github.com/johnbalvin/gobnb/search"
	"github.com/johnbalvin/gobnb/trace"
)

// optimized to work on this format: "https://www.airbnb.com/rooms/[roomID]"
func GetFromRoomURL(roomURL string, check_in_out search.Check, currency string, proxyURL *url.URL) (Data, error) {
	data, priceDependencyInput, cookies, err := getFromRoomURL(roomURL, proxyURL)
	if err != nil {
		return Data{}, trace.NewOrAdd(1, "main", "GetFromRoomURL", err, "")
	}
	if check_in_out.In.IsZero() || check_in_out.Out.IsZero() {
		return data, nil
	}
	price, err := priceDependencyInput.GetPrice(check_in_out, currency, cookies, proxyURL)
	if err != nil {
		return Data{}, trace.NewOrAdd(2, "main", "GetFromRoomURL", err, "")
	}
	data.Price = price
	return data, nil
}

func GetFromRoomID(roomID int64, check_in_out search.Check, currency string, proxyURL *url.URL) (Data, error) {
	roomURL := fmt.Sprintf("https://www.airbnb.com/rooms/%d", roomID)
	data, priceDependencyInput, cookies, err := getFromRoomURL(roomURL, proxyURL)
	if err != nil {
		return Data{}, trace.NewOrAdd(1, "main", "GetFromRoomID", err, "")
	}
	if check_in_out.In.IsZero() || check_in_out.Out.IsZero() {
		return data, nil
	}
	price, err := priceDependencyInput.GetPrice(check_in_out, currency, cookies, proxyURL)
	if err != nil {
		return Data{}, trace.NewOrAdd(2, "main", "GetFromRoomID", err, "")
	}
	data.Price = price
	return data, nil
}

func GetFromRoomIDAndDomain(roomID int64, domain string, check_in_out search.Check, currency string, proxyURL *url.URL) (Data, error) {
	roomURL := fmt.Sprintf("https://%s/rooms/%d", domain, roomID)
	data, priceDependencyInput, cookies, err := getFromRoomURL(roomURL, proxyURL)
	if err != nil {
		return Data{}, trace.NewOrAdd(1, "main", "GetFromRoomIDAndDomain", err, "")
	}
	if check_in_out.In.IsZero() || check_in_out.Out.IsZero() {
		return data, nil
	}
	price, err := priceDependencyInput.GetPrice(check_in_out, currency, cookies, proxyURL)
	if err != nil {
		return Data{}, trace.NewOrAdd(2, "main", "GetFromRoomIDAndDomain", err, "")
	}
	data.Price = price
	return data, nil
}
