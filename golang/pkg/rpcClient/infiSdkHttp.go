package rpcClient

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"golang/pkg/util"
	"log"
	"sort"
	"strconv"
	"strings"
)

var infiSdkHttpClientInstance *InfiSdkHttpClient

type InfiSdkHttpClient struct {
	appId   string //应用id
	signKey string //签名key
}

// Encrypt
func (_this *InfiSdkHttpClient) Encrypt(params map[string]string) (string, error) {
	if len(_this.signKey) == 0 {
		return "", fmt.Errorf("miss signKey")
	}
	keys := make([]string, 0, len(params))
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	keyValues := []string{}
	for _, k := range keys {
		if k == "signature" || len(k) == 0 {
			continue
		}
		keyValues = append(keyValues, k+"="+params[k])
	}
	p := strings.Join(keyValues, "&")
	mac := hmac.New(sha1.New, []byte(_this.signKey))
	mac.Write([]byte(p))
	var signature = fmt.Sprintf("%X", mac.Sum(nil))
	return signature, nil
}

// 计算api接口的签名
func (_this *InfiSdkHttpClient) CalculateWbsParams(params QueryParams) (QueryParams, string) {
	//在业务参数的基础上 补充签名参数
	params["appId"] = _this.appId
	params["expire"] = strconv.FormatInt(util.NowMs()+60*1000, 10) // 60秒内地址有效
	//params["expire"] = "1688970250420"
	businessParam := make(map[string]string)
	/*遍历params,将params中的kV值放入businessParam,并将value转为string*/
	for key, value := range params {
		businessParam[key] = fmt.Sprintf("%v", value)
	}
	// 业务参数和签名验证混合后排序
	var keys []string
	for key := range businessParam {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var res []string
	for _, key := range keys {
		if _, ok := businessParam[key]; !ok {
			continue
		}
		res = append(res, fmt.Sprintf("%s=%s", key, businessParam[key]))
	}

	// 排序后连接成字符串
	content := ""
	if len(res) > 0 {
		content = res[0]
		for i := 1; i < len(res); i++ {
			content += "&" + res[i]
		}
	}
	log.Printf("content: %s", content)
	h := hmac.New(sha1.New, []byte(_this.signKey))
	h.Write([]byte(content))
	params["signature"] = strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	return params, content + "&signature=" + params["signature"]
}

// 计算白板连接使用的签名
func (_this *InfiSdkHttpClient) CalculateBalanceParams(params QueryParams) (QueryParams, string) {
	//在业务参数的基础上 补充签名参数
	params["appId"] = _this.appId
	params["validBegin"] = strconv.FormatInt(util.NowS(), 10) // 60秒内地址有效
	params["validTime"] = strconv.FormatInt(120, 10)          // 60秒内地址有效

	businessParam := make(map[string]string)
	/*遍历params,将params中的kV值放入businessParam,并将value转为string*/
	for key, value := range params {
		businessParam[key] = fmt.Sprintf("%v", value)
	}
	// 业务参数和签名验证混合后排序
	var keys []string
	for key := range businessParam {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var res []string
	for _, key := range keys {
		if _, ok := businessParam[key]; !ok {
			continue
		}
		res = append(res, fmt.Sprintf("%s=%s", key, businessParam[key]))
	}

	// 排序后连接成字符串
	content := ""
	if len(res) > 0 {
		content = res[0]
		for i := 1; i < len(res); i++ {
			content += "&" + res[i]
		}
	}
	log.Printf("content: %s", content)
	h := hmac.New(sha1.New, []byte(_this.signKey))
	h.Write([]byte(content))
	params["signature"] = strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	return params, content + "&signature=" + params["signature"]
}

// 创建一个白板 "/u3wbs/wbs/nc/createBoard"
func (_this *InfiSdkHttpClient) CreateWhiteBoard(
	query CreateWhiteBoardQuery,
	params CreateWhiteBoardParams,
) *InfiSdkResponse[CreateWhiteBoardResponse] {
	queryParams, _ := _this.CalculateWbsParams(StructToQueryParams(query))
	res := HttpPost[InfiSdkResponse[CreateWhiteBoardResponse]](
		"https://api.infi.cn/u3wbs/wbs/nc/createBoard",
		queryParams,
		params,
	)
	return res
}

func InitInfiSdkHttpClient(appId, signKey string) *InfiSdkHttpClient {
	infiSdkHttpClientInstance = &InfiSdkHttpClient{
		appId:   appId,
		signKey: signKey,
	}
	return infiSdkHttpClientInstance
}
