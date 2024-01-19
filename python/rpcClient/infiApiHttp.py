import hashlib
import time
import requests
import hmac

class InfiApiHttpClient:
    def __init__(self, appId, signKey, infiWbsPath):
        self.appId = appId
        self.signKey = signKey
        self.infiWbsPath = infiWbsPath

    def calculate_wsb_params(self, params):
        # 在业务参数的基础上补充签名参数
        params['appId'] = self.appId
        params['expire'] = str((int(time.time()) + 60)*1000)  # 60秒内地址有效

        # 将参数按字典序排序
        sorted_params = sorted(params.items())

        # 连接参数字符串
        content = '&'.join(f'{key}={value}' for key, value in sorted_params)
        print(content)
        # 计算签名
        hmac_obj = hmac.new(self.signKey.encode(), content.encode(), hashlib.sha1)
        signature = hmac_obj.hexdigest().upper()
        params['signature'] = signature
        print(signature)
        return content + '&signature=' + signature

    def calculate_balance_params(self, params):
        # 在业务参数的基础上补充签名参数
        params['appId'] = self.appId
        params['validBegin'] = str((int(time.time())))  # 60秒内地址有效
        params['validTime'] = "120"  # 120秒内地址有效
        # 将参数按字典序排序
        sorted_params = sorted(params.items())

        # 连接参数字符串
        content = '&'.join(f'{key}={value}' for key, value in sorted_params)
        print(content)
        # 计算签名
        hmac_obj = hmac.new(self.signKey.encode(), content.encode(), hashlib.sha1)
        signature = hmac_obj.hexdigest().upper()
        params['signature'] = signature
        print(signature)
        return content + '&signature=' + signature

    def create_whiteboard(self, query, params):
        queryParams = self.calculate_wsb_params(query)
        url = f'{self.infiWbsPath}/u3wbs/wbs/nc/createBoard?'+queryParams
        print(url)
        try:
            response = requests.post(url, json=params)
            return response.json()
        except requests.exceptions.RequestException as error:
            print('Error creating whiteboard:', error)
            return None


