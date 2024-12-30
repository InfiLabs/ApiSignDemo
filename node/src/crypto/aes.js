const CryptoJS = require('crypto-js');
function genAes256(rawKey) {
    let keyString = rawKey
    const len = keyString.length;
    if (len > 32) {
        keyString = keyString.substring(0, 32);
    } else if (len < 32) {
        keyString += "xxxxxxxxxxxxxxxxxxxxxx".substring(0, 32 - len);
    }
    const key = CryptoJS.enc.Utf8.parse(keyString);
    const iv = CryptoJS.enc.Hex.parse("1934576290ABCBEF1264147890ACAE45");
    function encrypt(text) {
        const srcs = CryptoJS.enc.Utf8.parse(text);
        const encrypted = CryptoJS.AES.encrypt(srcs, key, {
            iv,
            mode: CryptoJS.mode.CBC,
            padding: CryptoJS.pad.Pkcs7,
        });
        return CryptoJS.enc.Base64.stringify(encrypted.ciphertext);
    }
    function decrypt(text) {
        const decryptResult = CryptoJS.AES.decrypt(text, key, {
            //  AES解密
            iv,
            mode: CryptoJS.mode.CBC,
            padding: CryptoJS.pad.Pkcs7,
        });
        const resData = decryptResult.toString(CryptoJS.enc.Utf8).toString();
        return resData;
    }
    return {
        encrypt,
        decrypt,
    };
}
module.exports = {
    genAes256
};
