/**
 * @author  zhaoliang.liang
 * @date  2024/1/19 0019 11:09
 */

package main

import (
	"fmt"
	"golang/accessorySDK"
	"golang/auth"
	"golang/define"
	"golang/pkg/crypto"
	"golang/pkg/utils"
	"golang/rpcClient"
	"log"
)

func CreateBoard() {
	infiHttpClient := rpcClient.GetInfiSdkHttpClient()
	// 创建一块白板
	res := infiHttpClient.CreateWhiteBoard(rpcClient.CreateWhiteBoardQuery{
		CreatorId: "infi",
	}, rpcClient.CreateWhiteBoardParams{})
	if res == nil {
		log.Printf("createBoard failed NetWorkError")
		return
	}
	if res.Code != 0 {
		log.Printf("createBoard failed httpError %s", res)
		return
	}
	log.Printf("createBoard success,response:%s", res)
	return
}

func CalculateBoardConnectParams() {
	// 开始生成query
	var infiBalanceQueryParams = rpcClient.InfiBalanceQueryParams{
		RecordId:       "recordId",      // 创建白板接口返回的recordId
		OwnerLoginName: "hostLoginName", // 白板创建者的用户ID,需要唯一
		LoginName:      "loginName",     // 白板连接者的用户ID,需要唯一
		UserName:       "userName",      // 白板连接者的用户名,用于显示在白板中
		UserType:       define.InfiUserTypeEditor,
		OpDays:         180,
		VersionDays:    180,
		Crypto:         1,
	}
	_, infiQueryUrl := rpcClient.GetInfiSdkHttpClient().
		CalculateBalanceParams(rpcClient.StructToQueryParams(infiBalanceQueryParams))
	log.Printf("infiQueryUrl:%s", infiQueryUrl)
}

func GetAccessoryToken() {
	sdkToken := accessorySDK.NewAccessorySDKToken()
	signKey := "1234567890123456" // 16/24/32位
	decrypt, err := crypto.CommonEncrypt(
		[]byte(sdkToken.String()),
		1,
		[]byte(signKey),
	)
	if err != nil {
		log.Fatal("encrypt failed", err)
	}
	log.Println(fmt.Sprintf("%s@%s", sdkToken.AppId, string(decrypt)))
}

func CreateAccessToken() {
	tokenObj := auth.CreateAccessToken("appId1", "recordId1", "loginName1", 3600)
	tokenStr, err := tokenObj.Build("appSecret")
	if err != nil {
		log.Fatal("token build failed", err)
	}
	log.Println(tokenStr)
	unPackTokenObj := auth.AccessToken{}
	if _, err := unPackTokenObj.Parse(tokenStr); err != nil {
		log.Fatal("token unpack failed", err)
	}
	log.Println(utils.ConvertToJsonStr(unPackTokenObj))
}

func main() {
	// 创建一块白板
	CreateBoard()
	// 计算白板连接参数
	CalculateBoardConnectParams()
	// 计算签到token
	GetAccessoryToken()
	// 计算accessToken
	CreateAccessToken()
}
