package org.InfiAccessorySdk.v1;

import cn.hutool.core.date.DateUtil;
import org.util.auth.InfiAccessorySdkToken;

public class Main {
    // https://developer.infi.cn/docs/guide/advancedFeatures/accessorySdk/intro/#%E5%87%86%E5%A4%87%E9%89%B4%E6%9D%83-token
    public static void InfiAccessorySdkApi_GetToken(String appId,String secret,String channelId,String loginName,String userType,long expire) {
        try {
            InfiAccessorySdkToken accessToken = new InfiAccessorySdkToken(appId, secret);
            String token = accessToken.generateToken(appId,channelId,loginName,userType,expire);
            System.out.println("Generated Access Token: " + token);
        } catch (Exception e) {
            System.out.println(e.getMessage());
        }
    }

    public static void main(String[] args) {
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
        String appId = "demo";
        String secret = "3oGPeNTjHdvxX2h7tR35OfPVOqYJQrzI";
        String channelId = "yourChannelId";
        String loginName = "yourLoginName";
        String userType = "host";
        long expire = 7*24*3600*1000+ DateUtil.current();
        InfiAccessorySdkApi_GetToken(appId,secret,channelId,loginName,userType,expire);

    }
}
