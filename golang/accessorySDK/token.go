package accessorySDK

import (
	"encoding/json"
	"time"
)

type InfiAccessoryUserType string

const (
	InfiAccessoryUserTypeHost  InfiAccessoryUserType = "host"  // 发起人
	InfiAccessoryUserTypeGuest InfiAccessoryUserType = "guest" // 用户
)

type InfiAccessorySDKToken struct {
	AppId     string                `json:"appId"`
	ChannelId string                `json:"channelId"`
	LoginName string                `json:"loginName"`
	UserType  InfiAccessoryUserType `json:"userType"`
	Expire    int64                 `json:"expire"`
}

// string
func (t *InfiAccessorySDKToken) String() string {
	jsonData, _ := json.Marshal(t)
	return string(jsonData)
}

func NewAccessorySDKToken() *InfiAccessorySDKToken {
	appId := "appId"
	loginName := "loginName"
	channelId := "channelId"
	userType := InfiAccessoryUserTypeHost
	// 设置过期时间24小时
	expire := time.Now().UnixMilli() + 24*3600*1000
	tokenObj := &InfiAccessorySDKToken{
		AppId:     appId,
		ChannelId: channelId,
		LoginName: loginName,
		UserType:  userType,
		Expire:    expire,
	}
	return tokenObj
}
