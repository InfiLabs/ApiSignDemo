const InfiApiHttpClient = require("../../../src/rpcClient/infiApiHttp");
const InfiUserType = require("../../../src/define/infiUserType");
async function createBoard(appId, secret) {
    let infiApiHttpClient = new InfiApiHttpClient(appId,secret, "https://api.infi.cn");
    let res = await infiApiHttpClient.createWhiteBoard({creatorId:"test"},{})
    if (!res) {
        throw new Error("create whiteboard failed");
    }
    console.log(res);
}

async function calculateBalanceParams(appId, secret){
    let infiApiHttpClient = new InfiApiHttpClient(appId,secret, "https://api.infi.cn");
    let boardQueryParams = {
        recordId:       "recordId",      // 创建白板接口返回的recordId
        ownerLoginName: "hostLoginName", // 白板创建者的用户ID,需要唯一
        loginName:      "loginName",     // 白板连接者的用户ID,需要唯一
        userName:       "userName",      // 白板连接者的用户名,用于显示在白板中
        userType:       InfiUserType.Editor, // 用户类型
        opDays:         180,
        versionDays:    180,
        crypto:         1,
    }
    let infiQuery = infiApiHttpClient.calculateBalanceParams(boardQueryParams)
    console.log(infiQuery)

}

async function main() {
    /*
     * title: 调用InfiApi创建一块画布
     * desc: 调用InfiWebSdk的RestApi创建一块画布并获取到recordId,用于提供给InfiWebSdk前端使用
     * url: https://developer.infi.cn/docs/restfulApi/board/create_board
     * @appId 为英飞分配的appId
     * @secret 为英飞分配的secret
     * */
    await createBoard("demo","3oGPeNTjHdvxX2h7tR35OfPVOqYJQrzI")

    /*
     * title: 计算 InfiWebSdk 所需的签名串
     * desc: 用于下发给前端InfiWebSdk中作为getQueryString参数使用
     * url: https://developer.infi.cn/docs/guide/prepare/getQueryString
     * @appId 为英飞分配的appId
     * @secret 为英飞分配的secret
     * */
    await calculateBalanceParams("demo","3oGPeNTjHdvxX2h7tR35OfPVOqYJQrzI")
}

main()
  .then((data) => {
    process.exit(0);
  })
  .catch((err) => {
    // eslint-disable-next-line no-console
    console.log(err);
    process.exit(1);
  });
