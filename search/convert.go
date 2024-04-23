package search

import (
	"log"
	"strconv"
	"strings"

	"github.com/johnbalvin/gobnb/trace"
	"github.com/johnbalvin/gobnb/utils"
)

func (rs rootdatapresentationstayssearchResults) standardize() []Data {
	var datas []Data
	for _, result := range rs.Results {
		lt := result.Listing
		pr := result.PricingQuote
		data := Data{
			RoomID:   result.Listing.Id,
			Category: result.Listing.RoomTypeCategory,
			Kind:     result.Listing.PdpUrlType,
			Name:     lt.Name,
			Title:    lt.Title,
			Type:     result.Listing.ListingObjType,
			Price: PriceData{
				Unit: UnitPrice{
					Qualifier: pr.StructuredStayDisplayPrice.PrimaryLine.Qualifier,
				},
			},
			Coordinates: Coordinates{
				Latitude: lt.Coordinate.Latitude,
				Longitud: lt.Coordinate.Longitude,
			},
		}
		for _, badge := range result.Listing.FormattedBadges {
			data.Badges = append(data.Badges, badge.LoggingContext.BadgeType)
		}
		splited := strings.Split(result.Listing.AvgRatingLocalized, " ")
		if len(splited) == 2 {
			rating, err := strconv.ParseFloat(splited[0], 32)
			if err != nil {
				errData := trace.NewOrAdd(1, "seach", "standardize", err, result.Listing.AvgRatingLocalized)
				log.Println(errData)
				continue
			}
			data.Rating.Value = float32(rating)
			reviewCountS := regexNumber.FindString(splited[1])
			reviewCount, err := strconv.Atoi(reviewCountS)
			if err != nil {
				errData := trace.NewOrAdd(2, "seach", "standardize", err, result.Listing.AvgRatingLocalized)
				log.Println(errData)
				continue
			}
			data.Rating.ReviewCount = reviewCount
		}
		priceToUse := pr.StructuredStayDisplayPrice.PrimaryLine.OriginalPrice
		if priceToUse == "" {
			priceToUse = pr.StructuredStayDisplayPrice.PrimaryLine.Price
		}
		if priceToUse != "" {
			ammount, currency, err := utils.ParsePriceSymbol(priceToUse)
			if err != nil {
				errData := trace.NewOrAdd(3, "seach", "standardize", err, pr.StructuredStayDisplayPrice.SecondaryLine.Price)
				log.Println(errData)
				continue
			}
			data.Price.Unit.CurrencySymbol = currency
			data.Price.Unit.Amount = ammount
		}
		if pr.StructuredStayDisplayPrice.PrimaryLine.DiscountedPrice != "" {
			ammount, _, err := utils.ParsePriceSymbol(pr.StructuredStayDisplayPrice.PrimaryLine.DiscountedPrice)
			if err != nil {
				errData := trace.NewOrAdd(4, "seach", "standardize", err, pr.StructuredStayDisplayPrice.SecondaryLine.Price)
				log.Println(errData)
				continue
			}
			data.Price.Unit.Discount = ammount
		}
		if pr.StructuredStayDisplayPrice.SecondaryLine.Price != "" {
			splited := strings.Split(pr.StructuredStayDisplayPrice.SecondaryLine.Price, " ")
			var priceToUse string
			switch len(splited) {
			case 2:
				priceToUse = splited[0]
			case 3:
				splited = splited[:len(splited)-1]
				priceToUse = strings.Join(splited, "")
			default:
				errData := trace.NewOrAdd(5, "seach", "standardize", trace.ErrParameter, pr.StructuredStayDisplayPrice.SecondaryLine.Price)
				log.Println(errData)
				continue
			}
			ammount, currency, err := utils.ParsePriceSymbol(priceToUse)
			if err != nil {
				errData := trace.NewOrAdd(6, "seach", "standardize", err, pr.StructuredStayDisplayPrice.SecondaryLine.Price)
				log.Println(errData)
				continue
			}
			data.Price.Total.CurrencySymbol = currency
			data.Price.Total.Amount = ammount
		}
		for _, priceDetail := range pr.StructuredStayDisplayPrice.ExplanationData.PriceDetails {
			for _, item := range priceDetail.Items {
				ammount, currency, err := utils.ParsePriceSymbol(item.PriceString)
				if err != nil {
					errData := trace.NewOrAdd(7, "seach", "standardize", err, item.PriceString)
					log.Println(errData)
					continue
				}
				data.Price.BreakDown = append(data.Price.BreakDown, PriceBreakDown{
					Description:    item.Description,
					Amount:         ammount,
					CurrencySymbol: currency,
				})
				switch item.DisplayComponentType {
				case "DISCOUNTED_EXPLANATION_LINE_ITEM":
					switch item.Description {
					case "Long stay discount":
						data.LongStayDiscount.Amount = ammount
						data.LongStayDiscount.CurrencySymbol = currency
					}
				case "DEFAULT_EXPLANATION_LINE_ITEM":
					switch item.Description {
					case "Cleaning fee":
						data.Fee.Cleaning.Amount = ammount
						data.Fee.Cleaning.CurrencySymbol = currency
					case "Airbnb service fee":
						data.Fee.Airbn.Amount = ammount
						data.Fee.Airbn.CurrencySymbol = currency
					}
				}
			}
		}
		for _, imgData := range lt.ContextualPictures {
			img := Img{
				URL: imgData.Picture,
			}
			data.Images = append(data.Images, img)
		}
		datas = append(datas, data)
	}
	return datas
}
