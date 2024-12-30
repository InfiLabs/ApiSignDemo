/**
 * @author  zhaoliang.liang
 * @date  2024/12/30 10:37
 */

package main

import (
	"golang/pkg/auth"
	"log"
)

func CreateAccessToken(appId, secret, recordId, loginName string, expire uint64) {
	// 注入appId和secret
	tokenObj := auth.CreateAccessToken(appId, recordId, loginName, expire)
	// 配置message权限,目前不需要配置
	//tokenObj.SetMessage(auth.MessageType(0), 0)
	// 生成token
	tokenStr, err := tokenObj.Build(secret)
	if err != nil {
		log.Fatal("token build failed", err)
	}
	log.Println(tokenStr)
}

func main() {
	/*
	 * title:计算 InfiWebSdk 所需的token
	 * desc: 用于下发给前端InfiWebSdk中作为initToken参数使用
	 * url: https://developer.infi.cn/docs/guide/prepare/initToken
	 * @appId 为英飞分配的appId
	 * @secret 为英飞分配的secret
	 * @recordId 为白板唯一ID,即上述接口创建获取到的recordId
	 * @loginName 为白板唯一用户名，为对接系统中的用户唯一标识，后续相关回调以及白板中确定用户身份都会用到
	 * @validTime 为签名有效期,单位为秒
	 * */
	appId := "1zDtJi"
	secret := "3oGPeNTjHdvxX2h7tR35OfPVOqYJQrzI"
	recordId := "recordId"
	loginName := "loginName"
	validTime := 3600
	CreateAccessToken(appId, secret, recordId, loginName, uint64(validTime))
}
