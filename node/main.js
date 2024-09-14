const InfiApiHttpClient = require("./rpcClient/infiApiHttp");
const config = require("./config/config");
const InfiUserType = require("./define/infiUserType");
async function createBoard(){
    let infiApiHttpClient = new InfiApiHttpClient(config.appId, config.signKey, config.infiWbsPath);
    let res = await infiApiHttpClient.createWhiteBoard({creatorId:"test"},{})
    if (!res) {
        throw new Error("create whiteboard failed");
    }
    console.log(res);
}

async function calculateBalanceParams(params){
    let infiApiHttpClient = new InfiApiHttpClient(config.appId, config.signKey, config.infiWbsPath);
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
    // 创建白板
    await createBoard()
    // 计算白板连接的签名
    //await calculateBalanceParams()
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
