package utils

import "time"

// NowS 当前时间戳（秒）
func NowS() int64 {
	return time.Now().Unix()
}

// NowMs 当前时间戳（毫秒）
func NowMs() int64 {
	return time.Now().UnixMilli()
}
