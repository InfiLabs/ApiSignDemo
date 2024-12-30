const crypto = require('crypto');
const { Buffer } = require('buffer');
const pack = require('../util/pack');
const url = require('url');

class AccessToken {
    constructor(appId, recordId, loginName, validTime) {
        this.appId = appId;
        this.recordId = recordId;
        this.loginName = loginName;
        this.validTime = validTime;
        this.signature = '';
        this.issueAt = Date.now();
        this.salt = this.random();
        this.message = {};
    }

    // Random number generation (max 99999999)
    random() {
        const maxVal = 99999999;
        return this.generateSecureRandomNumber(maxVal);
    }

    generateSecureRandomNumber(max) {
        return Math.floor(Math.random() * max);
    }

    // Signature Generation
    genSignature(appId, loginName, recordId, validTime, appSecret) {
        const params = {
            appId: appId,
            loginName: loginName,
            recordId: recordId,
            validTime: validTime.toString()
        };
        const signContent = this.generateQuery(false, params);
        const hmac = crypto.createHmac('sha1', appSecret);
        hmac.update(signContent);
        return hmac.digest('hex').toUpperCase();
    }

    // Setters for Salt and IssueAt
    setSalt() {
        this.salt = this.random();
    }

    setIssueAt() {
        this.issueAt =Date.now();
    }

    // Add message with key-value pair
    setMessage(key, value) {
        this.message[key] = value;
    }

    // Build the access token string
    build(appSecret) {
        let ret = '';

        this.signature = this.genSignature(this.appId, this.loginName, this.recordId, this.validTime, appSecret);

        const bufMsg = [];
        pack.packString(bufMsg, this.signature);
        pack.packUint64(bufMsg, this.salt);
        pack.packMapUint32(bufMsg, this.message);
        const bytesMsg = Buffer.concat(bufMsg);

        const bufContent = [];
        pack.packString(bufContent, this.recordId);
        pack.packUint64(bufContent, this.issueAt);
        pack.packString(bufContent, this.loginName);
        pack.packUint64(bufContent, this.validTime);
        pack.packString(bufContent, bytesMsg.toString('utf8'));

        const bytesContent = Buffer.concat(bufContent);

        ret = this.appId + '@' + bytesContent.toString('base64');
        return ret;
    }

    // Helper function to generate the query string
    generateQuery(needEncode, params) {
        const businessParam = {};
        Object.keys(params).forEach(key => {
            businessParam[key] = params[key].toString();
        });

        // Sorting the parameters by keys
        const keys = Object.keys(businessParam);
        keys.sort();

        let res = [];
        keys.forEach(key => {
            if (businessParam[key]) {
                if (needEncode) {
                    res.push(`${key}=${url.format({ pathname: businessParam[key] })}`);
                } else {
                    res.push(`${key}=${businessParam[key]}`);
                }
            }
        });

        // Combine the sorted parameters into a single query string
        return res.join('&');
    }
}

module.exports = AccessToken;
