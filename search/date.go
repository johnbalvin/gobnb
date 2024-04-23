package search

import (
	"fmt"
	"time"
)

func getStringDate(date time.Time) string {
	day := fmt.Sprintf("%d", date.Day())
	month := fmt.Sprintf("%d", date.Month())
	year := fmt.Sprintf("%d", date.Year())
	if len(day) == 1 {
		day = "0" + day
	}
	if len(month) == 1 {
		month = "0" + month
	}
	if len(year) == 1 {
		year = "0" + year
	}
	dateToUse := fmt.Sprintf("%s-%s-%s", year, month, day)
	return dateToUse
}
