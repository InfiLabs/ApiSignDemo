package auth

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"golang/pkg/utils"
	"log"
	"math/big"
	"net/url"
	"sort"
	"strings"
)

type Privileges uint16

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
	issueAt := uint64(utils.NowMs())
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
	_this.IssueAt = uint64(utils.NowMs())
}

func (_this *AccessToken) SetDefaultMessage() {
	_this.Message = make(map[uint16]uint32)
}

func (_this *AccessToken) Build(appSecret string) (string, error) {
	ret := ""

	_this.Signature = _this.GenSignature(_this.AppId, _this.LoginName, _this.RecordId, _this.ValidTime, appSecret)
	bufMsg := new(bytes.Buffer)
	if err := utils.PackString(bufMsg, _this.Signature); err != nil {
		log.Printf("packString(signature) %s err: %s\n", _this.Signature, err.Error())
		return ret, err
	}
	if err := utils.PackUint64(bufMsg, _this.Salt); err != nil {
		log.Printf("packUint64(salt) %d err: %s\n", _this.Salt, err.Error())
		return ret, err
	}
	if err := utils.PackMapUint32(bufMsg, _this.Message); err != nil {
		log.Printf("packMapUint32(message) %s err: %s\n", utils.ConvertToJsonStr(_this.Message), err.Error())
		return ret, err
	}
	bytesMsg := bufMsg.Bytes()

	bufContent := new(bytes.Buffer)
	if err := utils.PackString(bufContent, _this.RecordId); err != nil {
		log.Printf("packString(recordId) %s err: %s\n", _this.RecordId, err.Error())
		return ret, err
	}
	if err := utils.PackUint64(bufContent, _this.IssueAt); err != nil {
		log.Printf("packUint64(issueAt) %d err: %s\n", _this.IssueAt, err.Error())
		return ret, err
	}
	if err := utils.PackString(bufContent, _this.LoginName); err != nil {
		log.Printf("packString(loginName) %s err: %s\n", _this.LoginName, err.Error())
		return ret, err
	}
	if err := utils.PackUint64(bufContent, _this.ValidTime); err != nil {
		log.Printf("packUint32(crcRecordId) %d err: %s\n", _this.ValidTime, err.Error())
		return ret, err
	}
	if err := utils.PackString(bufContent, string(bytesMsg[:])); err != nil {
		log.Printf("packUint64(salt) %d err: %s\n", _this.Salt, err.Error())
		return ret, err
	}
	bytesContent := bufContent.Bytes()

	ret = _this.AppId + "@" + base64.StdEncoding.EncodeToString(bytesContent)
	return ret, nil
}

func (_this *AccessToken) Parse(originToken string) (res bool, err error) {
	defer PanicHandler()

	originStrList := strings.Split(originToken, "@")
	if len(originStrList) < 2 {
		return
	}

	var decodeByte []byte
	decodeByte, err = base64.StdEncoding.DecodeString(originToken[len(originStrList[0])+1:])
	recordId, issueAt, loginName, validTime, msgRawContent, err := UnPackContent(decodeByte)
	if err != nil {
		log.Printf("UnPackContent err: %s,can be ignored\n", err.Error())
		return
	}

	signature, salt, message, err := UnPackMessages(msgRawContent)
	if err != nil {
		log.Println("UnPackMessages err: ", err.Error())
		return
	}
	_this.AppId = originStrList[0]
	_this.Signature = signature
	_this.ValidTime = validTime
	_this.Salt = salt
	_this.IssueAt = issueAt
	_this.RecordId = recordId
	_this.LoginName = loginName
	_this.Message = message
	log.Println("content=", utils.ConvertToJsonStr(_this))
	return true, nil
}

func (_this *AccessToken) AddPrivilege(privilege Privileges, expireTimestamp uint32) {
	pri := uint16(privilege)
	_this.Message[pri] = expireTimestamp
}

func UnPackContent(buff []byte) (string, uint64, string, uint64, string, error) {
	in := bytes.NewReader(buff)
	recordId, err := utils.UnPackString(in)
	if err != nil {
		return "", 0, "", 0, "", err
	}

	issueAt, err := utils.UnPackUint64(in)
	if err != nil {
		return "", 0, "", 0, "", err
	}

	loginName, err := utils.UnPackString(in)
	if err != nil {
		return "", 0, "", 0, "", err
	}

	validTime, err := utils.UnPackUint64(in)
	if err != nil {
		return "", 0, "", 0, "", err
	}

	m, err := utils.UnPackString(in)
	if err != nil {
		return "", 0, "", 0, "", err
	}

	return recordId, issueAt, loginName, validTime, m, nil
}

func UnPackMessages(msgStr string) (string, uint64, map[uint16]uint32, error) {
	msgMap := make(map[uint16]uint32)

	msgByte := []byte(msgStr)
	in := bytes.NewReader(msgByte)

	signature, err := utils.UnPackString(in)
	if err != nil {
		return "", 0, msgMap, err
	}
	salt, err := utils.UnPackUint64(in)
	if err != nil {
		return "", 0, msgMap, err
	}

	length, err := utils.UnPackUint16(in)
	if err != nil {
		return "", 0, msgMap, err
	}
	for i := uint16(0); i < length; i++ {
		key, err := utils.UnPackUint16(in)
		if err != nil {
			return "", 0, msgMap, err
		}
		value, err := utils.UnPackUint32(in)
		if err != nil {
			return "", 0, msgMap, err
		}
		msgMap[key] = value
	}

	return signature, salt, msgMap, nil
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
