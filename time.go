package goutils

import "time"

func GetTimeStamp() int64 {

	return time.Now().Unix()

}

func GetTimeFormat(timeFormat string) (currentTime string) {

	return time.Now().Format(timeFormat)

}
