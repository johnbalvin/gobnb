package main

import (
	"net/url"

	"github.com/johnbalvin/gobnb/details"
	"github.com/johnbalvin/gobnb/search"
)

type Client struct {
	Currency string //ISO currency, example: USD, EUR
	ProxyURL *url.URL
}

func DefaultClient() Client {
	client := Client{
		Currency: "USD",
		ProxyURL: nil,
	}
	return client
}
func NewClient(currency string, proxyURL *url.URL) Client {
	client := Client{
		Currency: currency,
		ProxyURL: proxyURL,
	}
	return client
}

func (cl Client) DetailsFromRoomURL(roomURL string) (details.Data, error) {
	return details.GetFromRoomURL(roomURL, cl.Currency, cl.ProxyURL)
}
func (cl Client) DetailsFromRoomID(roomID int64) (details.Data, error) {
	return details.GetFromRoomID(roomID, cl.Currency, cl.ProxyURL)
}

func (cl Client) DetailsFromRoomIDAndDomain(roomID int64, domain string) (details.Data, error) {
	return details.GetFromRoomIDAndDomain(roomID, domain, cl.Currency, cl.ProxyURL)
}

func (cl Client) DetailsMainRoomIds(mailURL string) ([]int64, error) {
	return details.GetMainRoomIds(mailURL, cl.ProxyURL)
}

// coordinates should 2 points one from south and one from north(if you think it like a square, this points represent the two points of the diagonal from this square)
// zoom value from 1 - 20, so from the "square" like I said on the coorinates, this represents how much zoom on this squere there is
func (cl Client) SearchAll(zoomValue int, coordinates search.CoordinatesInput, check search.Check) ([]search.Data, error) {
	input := search.InputData{
		ZoomValue:   zoomValue,
		Coordinates: coordinates,
		Check:       check,
	}
	return input.SearchAll(cl.Currency, cl.ProxyURL)
}

// coordinates should 2 points one from south and one from north(if you think it like a square, this points represent the two points of the diagonal from this square)
// zoom value from 1 - 20, so from the "square" like I said on the coorinates, this represents how much zoom on this squere there is
func (cl Client) SearchFirstPage(zoomValue int, coordinates search.CoordinatesInput, check search.Check) ([]search.Data, error) {
	input := search.InputData{
		ZoomValue:   zoomValue,
		Coordinates: coordinates,
		Check:       check,
	}
	return input.SearchFirstPage(cl.Currency, cl.ProxyURL)
}
