/**
 * @author  zhaoliang.liang
 * @date  2024/1/19 0019 11:09
 */

package main

import (
	"golang/pkg/define"
	"golang/pkg/rpcClient"
	"log"
)

func CreateBoard(appId, secret string) {
	infiHttpClient := rpcClient.InitInfiSdkHttpClient(appId, secret)
	// 创建一块白板
	// creatorId 为白板创建者的用户ID,需要唯一
	res := infiHttpClient.CreateWhiteBoard(rpcClient.CreateWhiteBoardQuery{
		CreatorId: "creatorId",
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

func CalculateBoardConnectParams(appId, secret string) {
	infiHttpClient := rpcClient.InitInfiSdkHttpClient(appId, secret)
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
	_, infiQueryUrl := infiHttpClient.
		CalculateBalanceParams(rpcClient.StructToQueryParams(infiBalanceQueryParams))
	log.Printf("infiQueryUrl:%s", infiQueryUrl)
}

func main() {
	/*
	 * title: 调用InfiApi创建一块画布
	 * desc: 调用InfiWebSdk的RestApi创建一块画布并获取到recordId,用于提供给InfiWebSdk前端使用
	 * url: https://developer.infi.cn/docs/restfulApi/board/create_board
	 * @appId 为英飞分配的appId
	 * @secret 为英飞分配的secret
	 * */
	CreateBoard("demo", "3oGPeNTjHdvxX2h7tR35OfPVOqYJQrzI")

	/*
	 * title: 计算 InfiWebSdk 所需的签名串
	 * desc: 用于下发给前端InfiWebSdk中作为getQueryString参数使用
	 * url: https://developer.infi.cn/docs/guide/prepare/getQueryString
	 * @appId 为英飞分配的appId
	 * @secret 为英飞分配的secret
	 * */
	CalculateBoardConnectParams("demo", "3oGPeNTjHdvxX2h7tR35OfPVOqYJQrzI")
}
