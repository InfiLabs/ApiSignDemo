/**
 * @author  zhaoliang.liang
 * @date  2024/12/30 12:15
 */
const crypto = require("../../../src/crypto/aes");
async function getAccessoryToken(secret,sdkToken){
    let encrypt = await crypto.genAes256(secret).encrypt(JSON.stringify(sdkToken))
    console.log(`token:${sdkToken.appId}@${encrypt}`);

}

async function main() {
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
    let appId = "demo";
    let secret = "3oGPeNTjHdvxX2h7tR35OfPVOqYJQrzI";
    let channelId = "channelId";
    let loginName = "loginName";
    let userType = "host";
    let expire = 7*24*3600*1000 + Date.now();
    let sdkToken = {
        appId: appId,
        channelId: channelId,
        loginName: loginName,
        userType: userType,
        expire: expire
    }
    await getAccessoryToken(secret,sdkToken)
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
