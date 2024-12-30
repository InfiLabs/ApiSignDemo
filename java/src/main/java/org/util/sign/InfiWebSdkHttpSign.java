package org.util.sign;

import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;
import java.nio.charset.StandardCharsets;
import java.util.*;

public class InfiWebSdkHttpSign {
    private final String appId;
    private final String signKey;
    // 构造函数
    public InfiWebSdkHttpSign(String appId, String signKey) {
        this.appId = appId;
        this.signKey = signKey;
    }
    public String getInfiSign(Map<String,Object> params) throws Exception {
        // 业务参数，根据接口文档中定义，自行生成填写
        Map<String, Object> businessParam = new HashMap<>();
        businessParam.put("appId", this.appId);
        // 在业务参数的基础上补充签名参数
        businessParam.put("expire", System.currentTimeMillis()+60*1000); // 接口签名有效期默认60s
        // 将params循环放入业务参数中
        businessParam.putAll(params);

        // 业务参数和签名验证混合后排序
        List<String> afterSortParam = new ArrayList<>(businessParam.keySet());
        Collections.sort(afterSortParam);

        // 排序后连接成字符串
        StringBuilder content = new StringBuilder();
        for (String key : afterSortParam) {
            content.append(key).append("=").append(businessParam.get(key)).append("&");
        }
        content.deleteCharAt(content.length() - 1);

        // 使用分配的signKey来加密生成signature
        Mac hmacSha1 = Mac.getInstance("HmacSHA1");
        SecretKeySpec secretKey = new SecretKeySpec(this.signKey.getBytes(StandardCharsets.UTF_8), "HmacSHA1");
        hmacSha1.init(secretKey);
        byte[] signatureBytes = hmacSha1.doFinal(content.toString().getBytes(StandardCharsets.UTF_8));
        String signature = bytesToHex(signatureBytes).toUpperCase();

        // 生成的签名串signature和原来参数一起，得到最终的接口参数对象，发送给开放平台
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
