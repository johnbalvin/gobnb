package api

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/johnbalvin/gobnb/trace"
)

func Get(proxyURL *url.URL) (string, error) {
	req, err := http.NewRequest("GET", ep, nil)
	if err != nil {
		return "", trace.NewOrAdd(1, "main", "api", err, "")
	}
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("Accept-Language", "en")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Sec-Ch-Ua", `"Not_A Brand";v="8", "Chromium";v="120", "Google Chrome";v="120"`)
	req.Header.Add("Sec-Ch-Ua-Mobile", "?0")
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
		return "", trace.NewOrAdd(2, "main", "api", err, "")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", trace.NewOrAdd(3, "main", "api", err, "")
	}
	if resp.StatusCode != 200 {
		errData := fmt.Sprintf("status: %d headers: %+v", resp.StatusCode, resp.Header)
		return "", trace.NewOrAdd(4, "main", "api", trace.ErrStatusCode, errData)
	}
	apiKey := regxApiKey.FindString(string(body))
	apiKey = strings.ReplaceAll(apiKey, `"key":"`, "")
	apiKey = strings.ReplaceAll(apiKey, `"`, "")
	if apiKey == "" {
		return "", trace.NewOrAdd(5, "main", "api", trace.ErrEmpty, "")
	}
	return apiKey, nil
}
