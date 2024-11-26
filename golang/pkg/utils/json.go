package utils

import "encoding/json"

func ConvertToJsonStr(v interface{}) string {
	r, e := json.Marshal(v)
	if e != nil {
		return ""
	}
	return string(r)
}
