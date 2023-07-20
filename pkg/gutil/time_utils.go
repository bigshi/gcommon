package gutil

import (
	"time"
)

// GetTodayMinTimestamp
//
//	@Description: 今日零点毫秒
//	@return int64
func GetTodayMinTimestamp() int64 {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", timeStr)
	return t.UnixNano() / 1000000
}

// FormatDateTime
//
//	@Description: 格式化时间
//	@param t
//	@return string
func FormatDateTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// FormatDate
//
//	@Description: 格式化日期
//	@param t
//	@return string
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// FormatDate2
//
//	@Description: 格式化日期-简易
//	@param t
//	@return string
func FormatDate2(t time.Time) string {
	return t.Format("20060102")
}

// ParseDate
//
//	@Description: 解析成日期
//	@param timeStr
//	@return time.Time
func ParseDate(timeStr string) time.Time {
	if len(timeStr) > 10 {
		timeStr = timeStr[0:10]
	}
	loc, _ := time.LoadLocation("Local")
	date, _ := time.ParseInLocation("2006-01-02", timeStr, loc)
	return date
}

// ParseDate2
//
//	@Description: 解析成日期-简易
//	@param timeStr
//	@return time.Time
func ParseDate2(timeStr string) time.Time {
	if len(timeStr) > 8 {
		timeStr = timeStr[0:8]
	}
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation("20060102", timeStr, loc)
	return theTime
}

// ParseDateTime
//
//	@Description: 解析成时间
//	@param timeStr
//	@return time.Time
func ParseDateTime(timeStr string) time.Time {
	if len(timeStr) > 19 {
		timeStr = timeStr[0:19]
	}
	loc, _ := time.LoadLocation("Local")
	dateTime, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	return dateTime
}

// ToDateTime
//
//	@Description: 转时间戳
//	@param timestamp
//	@return time.Time
func ToDateTime(timestamp int64) time.Time {
	return time.Unix(timestamp/1000, (timestamp%1000)*1000000)
}
