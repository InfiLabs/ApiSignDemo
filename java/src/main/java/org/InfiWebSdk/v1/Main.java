package org.InfiWebSdk.v1;

import org.util.auth.InfiWebSdkBoardToken;

public class Main {
    public static void InfiWebSdkApi_GetBoardToken(String appId,String secret,String recordId,String loginName,long validTime) {
        try {
            // 注入appId和secret
            InfiWebSdkBoardToken accessToken = new InfiWebSdkBoardToken(appId, secret);
            // 配置message权限,目前不需要配置
            //accessToken.setMessage((short) 0,0);
            // 生成token
            String token = accessToken.generateToken(recordId,loginName,validTime);
            System.out.println("Generated Access Token: " + token);
        } catch (Exception e) {
            System.out.println(e.getMessage());
        }
    }


    public static void main(String[] args) {
        /*
         * title: 计算 InfiWebSdk 所需的token
         * desc: 用于下发给前端InfiWebSdk中作为initToken参数使用
         * url: https://developer.infi.cn/docs/guide/prepare/initToken
         * @appId 为英飞分配的appId
         * @secret 为英飞分配的secret
         * @recordId 为白板唯一ID,调用创建接口获取的recordId
         * @loginName 为白板唯一用户名，为对接系统中的用户唯一标识，后续相关回调以及白板中确定用户身份都会用到
         * @validTime 为签名有效期,单位为秒
         * */
        String appId = "1zDtJi";
        String secret = "3oGPeNTjHdvxX2h7tR35OfPVOqYJQrzI";
        String recordId = "recordId";
        String loginName = "loginName";
        long validTime = 3600L;
        InfiWebSdkApi_GetBoardToken(appId,secret,recordId,loginName,validTime);
    }
}
