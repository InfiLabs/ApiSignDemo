package org.util.auth;

import cn.hutool.core.date.DateUtil;

import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;
import java.io.ByteArrayOutputStream;
import java.io.DataOutputStream;
import java.nio.charset.StandardCharsets;
import java.security.SecureRandom;
import java.util.*;
import java.util.Base64;
import static org.util.pack.Pack.*;

public class InfiWebSdkBoardToken {
    private final String appId;         // 6位
    private final String secret;        // 32位
    private String recordId;      // 24位小写字母+数字
    private String loginName;     // 24位小写字母+数字
    private long validTime;       // 单位s
    private long issueAt;
    private long salt;
    private Map<Short, Integer> message;

    public InfiWebSdkBoardToken(String appId, String secret) {
        this.appId = appId;
        this.secret = secret;
        this.initSalt();
        this.initIssueAt();
        this.initMessage();
    }

    private long random() {
        long maxVal = 99999999L;
        return generateSecureRandomNumber(maxVal);
    }

    private long generateSecureRandomNumber(long max) {
        SecureRandom random = new SecureRandom();
        return Math.abs(random.nextLong() % max);
    }

    public String genSignature(String appSecret) {
        Map<String, String> params = new HashMap<>();
        params.put("appId", this.appId);
        params.put("loginName", this.loginName);
        params.put("recordId", this.recordId);
        params.put("validTime", String.valueOf(this.validTime));
        String signContent = generateQuery(false, params);
        return hmacSha1(signContent, appSecret);
    }

    private String hmacSha1(String data, String key) {
        try {
            Mac mac = Mac.getInstance("HmacSHA1");
            SecretKeySpec spec = new SecretKeySpec(key.getBytes(StandardCharsets.UTF_8), "HmacSHA1");
            mac.init(spec);
            byte[] bytes = mac.doFinal(data.getBytes(StandardCharsets.UTF_8));
            return bytesToHex(bytes);
        } catch (Exception e) {
            throw new RuntimeException("HMAC SHA1 error", e);
        }
    }

    private String bytesToHex(byte[] bytes) {
        StringBuilder hex = new StringBuilder();
        for (byte b : bytes) {
            hex.append(String.format("%02X", b));
        }
        return hex.toString();
    }

    private String generateQuery(boolean needEncode, Map<String, String> params) {
        List<String> keys = new ArrayList<>(params.keySet());
        Collections.sort(keys);

        List<String> res = new ArrayList<>();
        for (String key : keys) {
            String value = params.get(key);
            if (needEncode) {
                res.add(key + "=" + urlEncode(value));
            } else {
                res.add(key + "=" + value);
            }
        }

        return String.join("&", res);
    }

    private String urlEncode(String value) {
        try {
            return java.net.URLEncoder.encode(value, StandardCharsets.UTF_8);
        } catch (Exception e) {
            throw new RuntimeException("URL encoding error", e);
        }
    }

    public void initSalt() {
        this.salt = random();
    }

    public void initIssueAt() {
        this.issueAt = DateUtil.current();
    }

    public void initMessage() {
        this.message = new HashMap<>();
    }

    public void setMessage(short key, int value) {
        this.message.put(key, value);
    }

    public String generateToken(String recordId, String loginName, long validTime) throws Exception {
        this.recordId = recordId;
        this.loginName = loginName;
        this.validTime = validTime;
        String signature = genSignature(this.secret);

        ByteArrayOutputStream bufMsg = new ByteArrayOutputStream();
        try (DataOutputStream out = new DataOutputStream(bufMsg)) {
            packString(out, signature);
            packUint64(out, this.salt);
            packMapUint32(out, this.message);
        }

        ByteArrayOutputStream bufContent = new ByteArrayOutputStream();
        try (DataOutputStream out = new DataOutputStream(bufContent)) {
            packString(out, this.recordId);
            packUint64(out, this.issueAt);
            packString(out, this.loginName);
            packUint64(out, this.validTime);
            packString(out, bufMsg.toString());
        }

        byte[] bytesContent = bufContent.toByteArray();
        String base64EncodedContent = Base64.getEncoder().encodeToString(bytesContent);
        return this.appId + "@" + base64EncodedContent;
    }
}
