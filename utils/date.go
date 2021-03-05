package utils

import "time"

// 时间格式类型
type timeFormat string

const (
	// 年月日类型
	YYYYMMDD       timeFormat = "2006-01-02"
	YYYYMMDDHHIISS timeFormat = "2006-01-02 15:04:05"
)

// 快捷年月日类型获取
func Day() string {
	return DayFormat(YYYYMMDD)
}

func Second() string {
	return DayFormat(YYYYMMDDHHIISS)
}

func DayFormat(f timeFormat) string {
	return Format(time.Now(), f)
}

func Format(t time.Time, format timeFormat) string {
	return t.Format(string(format))
}
