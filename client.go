package gobnb

import "net/url"

type Client struct {
	Currency string //ISO currency, example: USD, EUR
	ProxyURL *url.URL
}

func DefaulClient() Client {
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

func (cl Client) GetFromRoomURL(roomURL string) (Data, error) {
	return GetFromRoomURL(roomURL, cl.Currency, cl.ProxyURL)
}
func (cl Client) GetFromRoomID(roomID string) (Data, error) {
	return GetFromRoomID(roomID, cl.Currency, cl.ProxyURL)
}

func (cl Client) GetFromRoomIDAndDomain(roomID, domain string) (Data, error) {
	return GetFromRoomIDAndDomain(roomID, domain, cl.Currency, cl.ProxyURL)
}

func (cl Client) GetMainRoomIds(mailURL string) ([]string, error) {
	return GetMainRoomIds(mailURL, cl.ProxyURL)
}
