/**
 * @author  zhaoliang.liang
 * @date  2024/12/30 12:15
 */
const AccessToken = require("../../../src/auth/infiWebSdkToken");

function createAccessToken(appId,secret,recordId,loginName,validTime){
    const token = new AccessToken(appId, recordId, loginName, validTime);
    const accessToken = token.build(secret);
    console.log("accessToken: ", accessToken);
}

async function main() {
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
    let appId = "demo"
    let secret = "3oGPeNTjHdvxX2h7tR35OfPVOqYJQrzI"
    let recordId = "recordId"
    let loginName = "loginName"
    let validTime = 3600
    await createAccessToken(appId,secret,recordId,loginName,validTime)
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
