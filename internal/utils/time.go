package utils

import (
	"fmt"
	"time"
)

// GetTodayMidnight 获取午夜0点的时间
func GetTodayMidnight() time.Time {
	now := time.Now().Local()
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return midnight
}

// GetYesterdayMidnight 获取昨天0点的时间
func GetYesterdayMidnight() time.Time {
	now := time.Now().Local()
	yesterdayMidnight := time.Date(now.Year(), now.Month(), now.Day()-1, 0, 0, 0, 0, now.Location())
	return yesterdayMidnight
}

// GetStartOfYear 获取新年第一天0点的时间
func GetStartOfYear() time.Time {
	now := time.Now().Local()
	startOfYear := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
	return startOfYear
}

// GetYearOf2000 2000年零点
func GetYearOf2000() time.Time {
	now := time.Now().Local()
	time := time.Date(2000, 1, 1, 0, 0, 0, 0, now.Location())
	return time
}

// IsTimeDifferenceThanNMinute 两个时间间隔是否超过n分钟,t1-t2
func IsTimeDifferenceThanNMinute(t1, t2 time.Time, n int64) bool {
	timeDifference := t1.Sub(t2)
	return timeDifference.Minutes() > 1
}

// GetCurrentDate 获取当前年月日
func GetCurrentDate() string {
	currentTime := time.Now()
	year, month, day := currentTime.Date()
	dateString := fmt.Sprintf("%d-%02d-%02d", year, month, day)
	return dateString
}
