package utils

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/johnbalvin/gobnb/trace"
)

var regxPrice = regexp.MustCompile(`\d+`)

func ParsePriceSymbol(priceRaw string) (float32, string, error) {
	priceRaw = strings.ReplaceAll(priceRaw, ",", "")
	priceNumber := regxPrice.FindString(priceRaw)
	priceCurrency := RemoveSpace(strings.ReplaceAll(priceRaw, priceNumber, ""))
	priceCurrency = strings.ReplaceAll(priceCurrency, "-", "")
	splited := strings.Split(priceRaw, "")
	if len(splited) < 2 {
		return 0, "", trace.NewOrAdd(1, "utils", "parsePriceSymbol", trace.ErrParameter, priceRaw)
	}
	price, err := strconv.ParseFloat(priceNumber, 32)
	if err != nil {
		return 0, "", trace.NewOrAdd(2, "utils", "parsePriceSymbol", err, priceRaw)
	}
	priceConverted := float32(price)
	if strings.HasPrefix(priceRaw, "-") {
		priceConverted *= -1
	}
	return priceConverted, priceCurrency, nil
}
