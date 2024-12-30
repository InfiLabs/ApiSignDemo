package accessorySDK

import (
	"encoding/json"
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

func NewAccessorySDKToken(appId string, loginName, channelId string, userType InfiAccessoryUserType, expire int64) *InfiAccessorySDKToken {
	tokenObj := &InfiAccessorySDKToken{
		AppId:     appId,
		ChannelId: channelId,
		LoginName: loginName,
		UserType:  userType,
		Expire:    expire,
	}
	return tokenObj
}
