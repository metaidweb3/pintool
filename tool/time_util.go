package tool

import (
	"time"
)

const hourTime = 3600 * 1000

var (
	l, _ = time.LoadLocation("UTC")
)

func MakeTimestamp() int64 {
	//time.Now().UnixMilli()
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func MakeDate(timestamp int64) string {
	timeFormat := "2006-01-02 15:04:05(UTC)"
	return time.Unix(timestamp/1000, 0).In(l).Format(timeFormat)
}

func MakeDateV2(timestamp int64) string {
	timeFormat := "20060102150405(UTC)"
	return time.Unix(timestamp/1000, 0).In(l).Format(timeFormat)
}

func MakeDateV3(timestamp int64) string {
	timeFormat := "2006-01-02"
	return time.Unix(timestamp/1000, 0).In(l).Format(timeFormat)
}

func Data2Timestamp(date string) int64 {
	theTime, err := time.ParseInLocation("2006-01-02 15:04:05(UTC)", date, l)
	if err != nil {
		return 0
	}
	timestamp := theTime.Unix() * 1000
	return timestamp
}

//00:00:00-time
func GetToday0Time() int64 {
	currentTime := time.Now()
	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, l)
	return startTime.UnixNano() / 1e6
}

//23:59:59-time
func GetToday24Time() int64 {
	currentTime := time.Now()
	endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, l)
	return endTime.UnixNano() / 1e6
}

//00:00:00-time
func GetYesterday0Time() int64 {
	currentTime := time.Now()
	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day()-1, 0, 0, 0, 0, l)
	return startTime.UnixNano() / 1e6
}

//23:59:59-time
func GetYesterday24Time() int64 {
	currentTime := time.Now()
	endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day()-1, 23, 59, 59, 0, l)
	return endTime.UnixNano() / 1e6
}

// Get 0 and 24 of the day according to the current time
func GetToday0And24Time() (int64, int64) {
	currentTime := time.Now()
	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, l)
	endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, l)
	return startTime.UnixNano() / 1e6, endTime.UnixNano() / 1e6
}

// Get 0 and 24 of the previous day according to the current time
func GetYesterday0And24Time() (int64, int64) {
	currentTime := time.Now()
	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day()-1, 0, 0, 0, 0, l)
	endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day()-1, 23, 59, 59, 0, l)
	return startTime.UnixNano() / 1e6, endTime.UnixNano() / 1e6
}
