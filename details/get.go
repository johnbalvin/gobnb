package details

import (
	"fmt"
	"net/url"

	"github.com/johnbalvin/gobnb/trace"
)

// optimized to work on this format: "https://www.airbnb.com/rooms/[roomID]"
func GetFromRoomURL(roomURL, currency string, proxyURL *url.URL) (Data, error) {
	data, priceDependencyInput, cookies, err := getFromRoomURL(roomURL, proxyURL)
	if err != nil {
		return Data{}, trace.NewOrAdd(1, "main", "GetFromRoomURL", err, "")
	}
	price, err := priceDependencyInput.GetPrice(currency, cookies, proxyURL)
	if err != nil {
		return Data{}, trace.NewOrAdd(2, "main", "GetFromRoomURL", err, "")
	}
	data.Price = price
	return data, nil
}

func GetFromRoomID(roomID int64, currency string, proxyURL *url.URL) (Data, error) {
	roomURL := fmt.Sprintf("https://www.airbnb.com/rooms/%d", roomID)
	data, priceDependencyInput, cookies, err := getFromRoomURL(roomURL, proxyURL)
	if err != nil {
		return Data{}, trace.NewOrAdd(1, "main", "GetFromRoomURL", err, "")
	}
	price, err := priceDependencyInput.GetPrice(currency, cookies, proxyURL)
	if err != nil {
		return Data{}, trace.NewOrAdd(2, "main", "GetFromRoomURL", err, "")
	}
	data.Price = price
	return data, nil
}

func GetFromRoomIDAndDomain(roomID int64, domain, currency string, proxyURL *url.URL) (Data, error) {
	roomURL := fmt.Sprintf("https://%s/rooms/%d", domain, roomID)
	data, priceDependencyInput, cookies, err := getFromRoomURL(roomURL, proxyURL)
	if err != nil {
		return Data{}, trace.NewOrAdd(1, "main", "GetFromRoomIDAndDomain", err, "")
	}
	price, err := priceDependencyInput.GetPrice(currency, cookies, proxyURL)
	if err != nil {
		return Data{}, trace.NewOrAdd(2, "main", "GetFromRoomIDAndDomain", err, "")
	}
	data.Price = price
	return data, nil
}
