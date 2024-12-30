/**
 * @author  zhaoliang.liang
 * @date  2024/12/30 10:41
 */

package main

import (
	"fmt"
	"golang/pkg/accessorySDK"
	"golang/pkg/crypto"
	"golang/pkg/util"
	"log"
)

func GetAccessoryToken(signKey string, sdkToken *accessorySDK.InfiAccessorySDKToken) {
	encrypt, err := crypto.CommonEncrypt(
		[]byte(sdkToken.String()),
		1,
		[]byte(signKey),
	)
	if err != nil {
		log.Fatal("encrypt failed", err)
	}
	log.Println(fmt.Sprintf("%s@%s", sdkToken.AppId, string(encrypt)))
}

func main() {
	/*
	 * title:计算 InfiAccessorySdk 所需的token
	 * desc: 用于下发给前端InfiAccessorySdk中作为initToken参数使用
	 * url: https://developer.infi.cn/docs/guide/advancedFeatures/accessorySdk/intro/#%E5%87%86%E5%A4%87%E9%89%B4%E6%9D%83-token
	 * @appId 为英飞分配的appId
	 * @secret 为英飞分配的secret
	 * @channelId 为频道ID
	 * @loginName 为白板唯一用户名，为对接系统中的用户唯一标识，后续相关回调以及白板中确定用户身份都会用到
	 * @userType 为用户类型 host/guest
	 * @expire 为token过期时间,单位为毫秒
	 * */
	appId := "demo"
	secret := "3oGPeNTjHdvxX2h7tR35OfPVOqYJQrzI"
	channelId := "yourChannelId"
	loginName := "yourLoginName"
	userType := accessorySDK.InfiAccessoryUserTypeHost
	expire := 7*24*3600*1000 + util.NowMs()
	sdkToken := accessorySDK.NewAccessorySDKToken(appId, loginName, channelId, userType, expire)
	GetAccessoryToken(secret, sdkToken)
}
