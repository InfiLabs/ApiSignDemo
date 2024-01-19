/**
 * @author  zhaoliang.liang
 * @date  2024/1/19 0019 11:17
 */

package util

import "time"

// NowMs 当前时间戳（毫秒）
func NowMs() int64 {
	return time.Now().UnixNano() / 1e6
}

// NowS 当前时间戳（秒）
func NowS() int64 {
	return time.Now().Unix()
}
