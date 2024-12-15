package utils

import (
	"time"
)

func String2Time(str string) (time.Time, error) {

	layout := "2006-01-02T15:04"
	parsedTime, err := time.Parse(layout, str)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}

func Time2String(t time.Time) string {
	layout := "2006-01-02T15:04"
	return t.Format(layout)
}

func DatetimeNow() time.Time {	
	t, _ := String2Time(time.Now().Format("2006-01-02T15:04"))
	return t
}