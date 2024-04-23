package utils

import (
	"bytes"
	"net/url"
	"regexp"
	"strings"
)

var regexSpace = regexp.MustCompile(`[\s ]+`)

func RemoveSpace(value string) string {
	return regexSpace.ReplaceAllString(strings.TrimSpace(value), " ")
}

func RemoveSpaceByte(value []byte) []byte {
	return regexSpace.ReplaceAll(bytes.TrimSpace(value), []byte(" "))
}

func ParseProxy(urlToParse, userName, password string) (*url.URL, error) {
	urlToUse, err := url.Parse(urlToParse)
	if err != nil {
		return nil, err
	}
	urlToUse.User = url.UserPassword(userName, password)
	return urlToUse, nil
}
