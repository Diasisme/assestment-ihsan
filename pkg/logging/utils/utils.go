package utils

import "time"

var (
	DATE_FORMAT = "2006-01-02"
)

func FormateDate(dateStr string) (date time.Time, err error) {
	date, err = time.Parse(DATE_FORMAT, dateStr)
	return
}

