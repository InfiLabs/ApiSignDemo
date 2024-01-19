package rpcClient

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"golang/util"
	"log"
	"reflect"
)

type QueryParams map[string]string
type BodyParams any

func StructToQueryParams(s interface{}) QueryParams {
	data := make(map[string]string)

	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		// 忽略未导出的字段和非字符串类型的 tag
		if field.PkgPath != "" {
			continue
		}

		if field.Type.Kind() == reflect.Struct {
			// 递归处理嵌套的结构体
			for k, v := range StructToQueryParams(value) {
				data[k] = v
			}
			continue
		}

		tag := field.Tag.Get("json")
		if tag == "" {
			tag = field.Name
		}

		data[tag] = fmt.Sprintf("%v", value)
	}

	return data
}

func HttpPost[T any](url string, query QueryParams, bodyParams BodyParams) *T {
	client := resty.New()
	reqClient := client.R()
	if query != nil {
		reqClient.SetQueryParams(query)
	}
	if bodyParams != nil {
		log.Printf(
			"url:%s , BodyParams:%s\n",
			reqClient.URL,
			util.MustJsonMarshalByte(bodyParams, `{}`),
		)
		reqClient.SetBody(bodyParams)
	}
	resp, err := reqClient.Post(url)

	if err != nil {
		log.Printf("url:%s , resp Error: %s\n", reqClient.URL, err.Error())
		return nil
	}

	log.Printf("url:%s , Response Body:%s\n", reqClient.URL, string(resp.Body()))
	// 对body进行解析
	var response T
	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		log.Printf("json Error: %s\n", err.Error())
		return nil
	}
	return &response
}
