package goutils

import (
	"fmt"
	"time"
)

func GetTimeStamp() int64 {

	return time.Now().Unix()

}

func GetTimeFormat(timeFormat string) (currentTime string) {

	return time.Now().Format(timeFormat)

}

func GetDurationDesc(durationTime time.Duration) (desc string) {

	duration := durationTime.Seconds()

	switch {
	case duration <= 60*10:
		desc = "刚刚"
	case duration <= 60*60:
		desc = fmt.Sprintf("%d分钟前", int(duration)/60)
	case duration <= 60*60*24:
		desc = fmt.Sprintf("%d小时前", int(duration)/60/60)
	case duration <= 60*60*24*30:
		desc = fmt.Sprintf("%d天前", int(duration)/3600/24)
	default:
		desc = fmt.Sprintf("%d月前", int(duration)/3600/24/30)
	}

	return desc

}
