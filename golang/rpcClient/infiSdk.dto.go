package rpcClient

import (
	"golang/define"
	"golang/util"
)

type InfiSdkResponse[T any] struct {
	Code  int64  `json:"code"`
	ReqId string `json:"reqId"`
	Obj   T      `json:"obj"`
}

func (s InfiSdkResponse[any]) String() string {
	return util.Prettify(s)
}

type CreateWhiteBoardQuery struct {
	// 白板创建者的用户ID
	CreatorId string `json:"creatorId"`
}
type CreateWhiteBoardParams struct {
}

type CreateWhiteBoardResponse struct {
	RecordId string `json:"recordId"`
	BoardId  string `json:"boardId"`
}

// 白板连接query参数
type InfiBalanceQueryParams struct {
	RecordId       string              `form:"recordId"       json:"recordId"       binding:"required"`
	OwnerLoginName string              `form:"ownerLoginName" json:"ownerLoginName" binding:"required"`
	LoginName      string              `form:"loginName"      json:"loginName"      binding:"required"`
	UserName       string              `form:"userName"       json:"userName"       binding:"required"`
	UserType       define.InfiUserType `form:"userType"       json:"userType"       binding:"required"`
	OpDays         int                 `form:"opDays"         json:"opDays"         binding:"required"`
	VersionDays    int                 `form:"versionDays"    json:"versionDays"    binding:"required"`
	Crypto         int                 `form:"crypto"         json:"crypto"         binding:"required"`
}
