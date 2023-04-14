package util

import (
	"fmt"
	"time"
)

//
// GetNowUnixMills
//  @Description: 获取当前时间戳-毫秒
//  @return int64
//
func GetNowUnixMills() int64 {
	return time.Now().UnixNano() / 1e6
}

//
// ToDateTimeStr
//  @Description: 格式化时间
//  @param t
//  @return string
//
func ToDateTimeStr(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

//
// ToDateStr
//  @Description: 格式化日期
//  @param t
//  @return string
//
func ToDateStr(t time.Time) string {
	return t.Format("2006-01-02")
}

//
// ToDateStrOfSample
//  @Description: 格式化日期-简易
//  @param t
//  @return string
//
func ToDateStrOfSample(t time.Time) string {
	return t.Format("20060102")
}

//
// FromStrToDateOfSample
//  @Description: 获取日期-简易
//  @param timeStr
//  @return time.Time
//
func FromStrToDateOfSample(timeStr string) time.Time {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation("20060102", timeStr, loc)
	return theTime
}

//
// FromStrToDateTime
//  @Description: 获取时间
//  @param timeStr
//  @return time.Time
//
func FromStrToDateTime(timeStr string) time.Time {
	if len(timeStr) > 19 {
		timeStr = timeStr[0:19]
	}
	loc, _ := time.LoadLocation("Local")
	dateTime, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	return dateTime
}

//
// FromStrToDate
//  @Description: 获取日期
//  @param timeStr
//  @return time.Time
//
func FromStrToDate(timeStr string) time.Time {
	if len(timeStr) > 10 {
		timeStr = timeStr[0:10]
	}
	loc, _ := time.LoadLocation("Local")
	date, _ := time.ParseInLocation("2006-01-02", timeStr, loc)
	return date
}

//
// FromTimestampToDateTime
//  @Description: 转时间戳
//  @param timestamp
//  @return time.Time
//
func FromTimestampToDateTime(timestamp int64) time.Time {
	return time.Unix(timestamp/1000, (timestamp%1000)*1000000)
}

//
// ToTimestamp
//  @Description: 转毫秒
//  @param t
//  @return int64
//
func ToTimestamp(t time.Time) int64 {
	return t.UnixNano() / 1000000
}

//
// GetTodayMinTimestamp
//  @Description: 今日零点毫秒
//  @return int64
//
func GetTodayMinTimestamp() int64 {
	timeStr := time.Now().Format("2006-01-02")
	fmt.Println("timeStr:", timeStr)
	t, _ := time.Parse("2006-01-02", timeStr)
	return t.UnixNano() / 1000000
}
