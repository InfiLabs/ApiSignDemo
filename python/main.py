# createBoard
from rpcClient.infiApiHttp import InfiApiHttpClient
from config.config import AppId, SignKey, InfiWbsPath


def create_board():
    infiApiClient = InfiApiHttpClient(AppId, SignKey, InfiWbsPath)

    query = {
        # query参数
        "creatorId": 'infi',
    }

    params = {
        # Fill in the request body parameters
    }

    response = infiApiClient.create_whiteboard(query, params)
    print(response)


def calculate_board_query_sign():
    infiApiClient = InfiApiHttpClient(AppId, SignKey, InfiWbsPath)
    from define.infiUserType import InfiUserType
    query = {
        # query参数
        "recordId": "recordId",  # 创建白板接口返回的recordId
        "ownerLoginName": "hostLoginName",  # 白板创建者的用户ID,需要唯一
        "loginName": "loginName",  # 白板连接者的用户ID,需要唯一
        "userName": "userName",  # 白板连接者的用户名,用于显示在白板中
        "userType": InfiUserType.Editor.value,  # 用户类型
        "opDays": 180,
        "versionDays": 180,
        "crypto": 1,
    }

    response = infiApiClient.calculate_balance_params(query)
    print(response)


if __name__ == '__main__':
    # 创建白板
    create_board()
    # 计算白板连接签名
    calculate_board_query_sign()
