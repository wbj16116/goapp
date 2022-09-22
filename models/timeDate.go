package models

import (
	"time"
	"strconv"
)

func UnixToDate(timestamp int64) string {

	t := time.Unix(int64(timestamp), 0)

	return t.Format("2006-01-02 15:04:05")
}

func GetStringTime() string {
	return strconv.FormatInt(time.Now().Unix(), 10)  //字符串类型时间戳 strconv.FormatInt(int 转字符串)
													//strconv.Itoa(字符串转int)
}

func GetIntTime() int64 {
	return time.Now().Unix()   //int64类型时间戳
}
