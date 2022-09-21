package utils

import "time"

func GetNowTime() string {
	t := time.Now().Format("2006-01-02 15:04:05")
	return t
}

// StrToTime 字符串转time
func StrToTime(str string) time.Time {
	p, _ := time.Parse("2006-01-02 15:04:05", str)
	return p
}

// TimeToStr 日期转字符串
func TimeToStr(t time.Time) string {
	return t.Format("2006-01-02 03:04:05")
}
