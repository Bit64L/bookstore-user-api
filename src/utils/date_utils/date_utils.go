package date_utils

import "time"

var(
	//"Mon Jan 2 15:04:05 -0700 MST 2006"
	apiDateLayout="2006-01-02T15:04:05Z"
)


func GetNow() time.Time{
	return time.Now().UTC()
}

func GetNowString() string{
	return GetNow().Format(apiDateLayout)
}

