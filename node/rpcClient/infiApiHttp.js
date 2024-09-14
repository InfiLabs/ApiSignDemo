const crypto = require('crypto');
const axios = require('axios');
class InfiApiHttpClient {
    constructor(appId,signKey,infiWbsPath) {
        this.appId = appId;
        this.signKey = signKey;
        this.infiWbsPath = infiWbsPath;
    }
    // 计算api请求的签名
    calculateWbsParams(params) {
        // 在业务参数的基础上补充签名参数
        params.appId = this.appId;
        params.expire = (Date.now() + 60 * 1000).toString(); // 60秒内地址有效

        // 将参数转换为字符串
        const businessParam = {};
        for (const [key, value] of Object.entries(params)) {
            businessParam[key] = String(value);
        }

        // 排序参数
        const keys = Object.keys(businessParam).sort();

        const res = [];
        for (const key of keys) {
            if (businessParam.hasOwnProperty(key)) {
                res.push(`${key}=${businessParam[key]}`);
            }
        }

        // 连接参数字符串
        let content = '';
        if (res.length > 0) {
            content = res.join('&');
        }

        console.log(`content: ${content}`);

        // 计算签名
        const hmac = crypto.createHmac('sha1', this.signKey);
        hmac.update(content);
        const signature = hmac.digest('hex').toUpperCase();

        params.signature = signature;
        console.log(`signature: ${signature}`);
        return `${content}&signature=${signature}`;
    }
    // 计算白板连接的签名
    calculateBalanceParams(params) {
        // 在业务参数的基础上补充签名参数
        params.appId = this.appId;
        params.validBegin = Math.floor(Date.now() / 1000).toString();
        params.validTime = '120'; // 120秒内地址有效

        // 将参数转换为字符串
        const businessParam = {};
        for (const [key, value] of Object.entries(params)) {
            businessParam[key] = String(value);
        }

        // 排序参数
        const keys = Object.keys(businessParam).sort();

        const res = [];
        for (const key of keys) {
            if (businessParam.hasOwnProperty(key)) {
                res.push(`${key}=${businessParam[key]}`);
            }
        }

        // 连接参数字符串
        let content = '';
        if (res.length > 0) {
            content = res.join('&');
        }

        console.log(`content: ${content}`);

        // 计算签名
        const hmac = crypto.createHmac('sha1', this.signKey);
        hmac.update(content);
        const signature = hmac.digest('hex').toUpperCase();

        params.signature = signature;
        return `${content}&signature=${signature}`;
    }
    async createWhiteBoard(query, params) {
        const queryParams = this.calculateWbsParams(query);
        console.log(queryParams);
        const url = this.infiWbsPath + '/u3wbs/wbs/nc/createBoard?'+queryParams;

        try {
            const response = await axios.post(url, params);
            return response.data;
        } catch (error) {
            console.error('Error creating whiteboard:', error);
            return null;
        }
    }

}

module.exports = InfiApiHttpClient;
