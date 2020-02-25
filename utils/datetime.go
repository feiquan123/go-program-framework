package utils

import (
	"strings"
	"time"
)

// 时间格式
const (
	TimeFormat1      = "2006-01-02T15:04:05+08:00"
	TimeFormat2      = "2006-01-02 15:04:05 -0700 MST"
	TimeFormat3      = "2006-01-02 15:04:05"
	TimeFormat4      = "2006-01-02 15:04:05.999"
	TimeFormat5      = "2006-01-02"
	TimeFormatSuffix = " +0800 CST"
)

var (
	// ShortMonth : 月份短昵称
	ShortMonth = map[string]time.Month{
		"Jan": time.January,
		"Feb": time.February,
		"Mar": time.March,
		"Apr": time.April,
		"May": time.May,
		"Jun": time.June,
		"Jul": time.July,
		"Aug": time.August,
		"Sep": time.September,
		"Oct": time.October,
		"Nov": time.November,
		"Dec": time.December,
	}
)

// ConvertTimestampUnix : convart timestamp to unix nano
func ConvertTimestampUnix(timeformat, timestampStr string) (timestamp int64, err error) {
	ctime, err := time.ParseInLocation(timeformat, timestampStr, time.Local)
	if err != nil {
		return 0, err
	}
	return ctime.Unix(), nil
}

// ConvertTimestampUnixNano : convert timstamp to unix nano
func ConvertTimestampUnixNano(timeformat, timestampStr string) (timestamp int64, err error) {
	ctime, err := time.ParseInLocation(timeformat, timestampStr, time.Local)
	if err != nil {
		return 0, err
	}
	return ctime.UnixNano(), nil
}

// TimeStamp2String 2020-01-10_09-35-20
func TimeStamp2String(timeFormat string, timstamp int64) string {
	strTime := time.Unix(timstamp, 0).Format(timeFormat)
	strTime = strings.Replace(strTime, " ", "_", -1)
	strTime = strings.Replace(strTime, ":", "-", -1)
	return strTime
}
