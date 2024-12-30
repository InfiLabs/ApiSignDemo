package org.util.auth;

import cn.hutool.json.JSONObject;
import cn.hutool.json.JSONUtil;
import org.util.crypto.AesCrypto;

public class InfiAccessorySdkToken {
    private final String appId; // appId
    private final String secret; // secret

    // 构造函数，传入 appId 和 secretKey
    public InfiAccessorySdkToken(String appId, String secret) {
        this.appId = appId;
        this.secret = secret;
    }

    public String generateToken(String appId,String channelId,String loginName,String userType,long expire) throws Exception {
        JSONObject params = JSONUtil.createObj()
                .put("appId", appId)
                .put("channelId", channelId)
                .put("loginName", loginName)
                .put("userType", userType)
                .put("expire", expire);
        AesCrypto aesCrypto = new AesCrypto();
        String encryptStr = aesCrypto.cbcEncrypt(params.toString(), this.secret);
        return this.appId + "@" + encryptStr;

    }
}
