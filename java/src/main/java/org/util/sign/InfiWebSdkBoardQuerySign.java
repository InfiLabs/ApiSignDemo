package org.util.sign;

import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;
import java.nio.charset.StandardCharsets;
import java.util.*;


public class InfiWebSdkBoardQuerySign {
    private final String appId;
    private final String signKey;
    public InfiWebSdkBoardQuerySign(String appId, String signKey) {
        this.appId = appId;
        this.signKey = signKey;
    }

    public String getInfiBoardQuerySign(Map<String,Object> params) throws Exception {
        Map<String, Object> businessParam = new HashMap<>();
        businessParam.put("appId", this.appId);
        businessParam.put("validBegin", System.currentTimeMillis()/1000); // 签名有效期开始时间，单位为秒
        businessParam.put("validTime", "120"); // 签名有效期,如果业务上不想重复下发签名，可以设置大一点
        businessParam.putAll(params);
        List<String> afterSortParam = new ArrayList<>(businessParam.keySet());
        Collections.sort(afterSortParam);
        StringBuilder content = new StringBuilder();
        for (String key : afterSortParam) {
            content.append(key).append("=").append(businessParam.get(key)).append("&");
        }
        content.deleteCharAt(content.length() - 1);
        Mac hmacSha1 = Mac.getInstance("HmacSHA1");
        SecretKeySpec secretKey = new SecretKeySpec(this.signKey.getBytes(StandardCharsets.UTF_8), "HmacSHA1");
        hmacSha1.init(secretKey);
        byte[] signatureBytes = hmacSha1.doFinal(content.toString().getBytes(StandardCharsets.UTF_8));
        String signature = bytesToHex(signatureBytes).toUpperCase();
        businessParam.put("signature", signature);
        return content+"&signature="+signature;
    }

    // 将字节数组转换为十六进制字符串
    private static String bytesToHex(byte[] bytes) {
        StringBuilder hexString = new StringBuilder();
        for (byte b : bytes) {
            String hex = Integer.toHexString(0xFF & b);
            if (hex.length() == 1) {
                hexString.append('0');
            }
            hexString.append(hex);
        }
        return hexString.toString();
    }
}
