/**
 * @author  zhaoliang.liang
 * @date  2024/1/19 0019 11:14
 */

package util

import "encoding/json"

func MustJsonMarshalByte(v interface{}, defaultValue string) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		return []byte(defaultValue)
	}
	return b
}

func Prettify(i interface{}) string {
	resp, _ := json.Marshal(i)
	return string(resp)
}

func ConvertToJsonStr(v interface{}) string {
	r, e := json.Marshal(v)
	if e != nil {
		return ""
	}
	return string(r)
}
