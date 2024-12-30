package auth

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"golang/pkg/pack"
	"golang/pkg/util"
	"log"
	"math/big"
	"net/url"
	"sort"
)

type MessageType uint16

// AccessToken 供外部对接app对接用的
type AccessToken struct {
	AppId     string            `json:"appId"     binding:"required"` // 6位
	RecordId  string            `json:"recordId"  binding:"required"` // 24位小写字母+数字
	LoginName string            `json:"loginName" binding:"required"` // 24位小写字母+数字
	ValidTime uint64            `json:"validTime" binding:"required"` // 单位s
	Signature string            `json:"signature"`
	IssueAt   uint64            `json:"issueAt"`
	Salt      uint64            `json:"salt"`
	Message   map[uint16]uint32 `json:"message"`
}

func random() uint64 {
	maxVal := int64(99999999)
	val, err := generateSecureRandomNumber(maxVal)
	if err != nil {
		log.Fatalln("generateSecureRandomNumber err ", err)
		return 1
	}
	return val
}

func generateSecureRandomNumber(max int64) (uint64, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return 0, err
	}
	return n.Uint64(), nil
}

func PanicHandler() {
	if r := recover(); r != nil {
		log.Fatalln("error: ", r)
	}
}

func CreateAccessToken(appId, recordId, loginName string, validTime uint64) *AccessToken {
	issueAt := uint64(util.NowMs())
	salt := random()
	message := make(map[uint16]uint32)

	return &AccessToken{
		AppId:     appId,
		RecordId:  recordId,
		LoginName: loginName,
		ValidTime: validTime,
		IssueAt:   issueAt,
		Salt:      salt,
		Message:   message,
	}
}

func (_this *AccessToken) GenSignature(appId, loginName, recordId string, validTime uint64, appSecret string) string {
	params := map[string]string{
		"appId":     appId,
		"loginName": loginName,
		"recordId":  recordId,
		"validTime": fmt.Sprintf("%d", validTime),
	}
	signContent := GenerateQuery(false, params)
	h := hmac.New(sha1.New, []byte(appSecret))
	h.Write([]byte(signContent))
	signature := fmt.Sprintf("%X", h.Sum(nil))
	return signature
}

func (_this *AccessToken) SetSalt() {
	_this.Salt = random()
}

func (_this *AccessToken) SetIssueAt() {
	_this.IssueAt = uint64(util.NowMs())
}

func (_this *AccessToken) SetMessage(key MessageType, value uint32) {
	_this.Message[uint16(key)] = value
}

func (_this *AccessToken) Build(appSecret string) (string, error) {
	ret := ""

	_this.Signature = _this.GenSignature(_this.AppId, _this.LoginName, _this.RecordId, _this.ValidTime, appSecret)
	bufMsg := new(bytes.Buffer)
	if err := pack.PackString(bufMsg, _this.Signature); err != nil {
		log.Printf("packString(signature) %s err: %s\n", _this.Signature, err.Error())
		return ret, err
	}
	if err := pack.PackUint64(bufMsg, _this.Salt); err != nil {
		log.Printf("packUint64(salt) %d err: %s\n", _this.Salt, err.Error())
		return ret, err
	}
	if err := pack.PackMapUint32(bufMsg, _this.Message); err != nil {
		log.Printf("packMapUint32(message) %s err: %s\n", util.ConvertToJsonStr(_this.Message), err.Error())
		return ret, err
	}
	bytesMsg := bufMsg.Bytes()

	bufContent := new(bytes.Buffer)
	if err := pack.PackString(bufContent, _this.RecordId); err != nil {
		log.Printf("packString(recordId) %s err: %s\n", _this.RecordId, err.Error())
		return ret, err
	}
	if err := pack.PackUint64(bufContent, _this.IssueAt); err != nil {
		log.Printf("packUint64(issueAt) %d err: %s\n", _this.IssueAt, err.Error())
		return ret, err
	}
	if err := pack.PackString(bufContent, _this.LoginName); err != nil {
		log.Printf("packString(loginName) %s err: %s\n", _this.LoginName, err.Error())
		return ret, err
	}
	if err := pack.PackUint64(bufContent, _this.ValidTime); err != nil {
		log.Printf("PackUint64(ValidTime) %d err: %s\n", _this.ValidTime, err.Error())
		return ret, err
	}
	if err := pack.PackString(bufContent, string(bytesMsg[:])); err != nil {
		log.Printf("PackString(bytesMsg) %d err: %s\n", string(bytesMsg[:]), err.Error())
		return ret, err
	}
	bytesContent := bufContent.Bytes()

	ret = _this.AppId + "@" + base64.StdEncoding.EncodeToString(bytesContent)
	return ret, nil
}

func GenerateQuery(needEncode bool, params map[string]string) string {
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
		if needEncode {
			res = append(res, fmt.Sprintf("%s=%s", key, url.QueryEscape(businessParam[key])))
		} else {
			res = append(res, fmt.Sprintf("%s=%s", key, businessParam[key]))
		}
	}
	// 排序后连接成字符串
	content := ""
	if len(res) > 0 {
		content = res[0]
		for i := 1; i < len(res); i++ {
			content += "&" + res[i]
		}
	}
	return content
}
