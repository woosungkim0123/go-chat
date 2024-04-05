package util

import "time"

func GetCurrentDateWithFormat() string {
	location, _ := time.LoadLocation("Asia/Seoul")
	currentTime := time.Now().UTC().In(location)
	return currentTime.Format("1/02 15:04:05")
}

func GetCurrentDate() time.Time {
	location, _ := time.LoadLocation("Asia/Seoul")
	return time.Now().UTC().In(location)
}

func ConvertDateToString(date time.Time) string {
	return date.Format("1/02 15:04:05")
}
