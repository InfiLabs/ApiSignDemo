package org.example;


import java.util.HashMap;
import cn.hutool.http.HttpUtil;

//TIP 要<b>运行</b>代码，请按 <shortcut actionId="Run"/> 或
// 点击间距中的 <icon src="AllIcons.Actions.Execute"/> 图标。
public class Main {
    public static void main(String[] args) {
        try {
            /*
            *
            *
            * ##########################################以下为生成调用英飞API所需用的sign############################################
            *
            * */
            // 初始化签名对象
            InfiHttpSignDemo infiHttpSignDemo = new InfiHttpSignDemo("appId","signkey");
            HashMap<String, Object> params = new HashMap<String,Object>();
            // 此处构造接口文档所需的参数
            // 除appId expire signature 这些参数由sign函数自行注入
            params.put("creatorId","test");
            String sign = infiHttpSignDemo.getInfiSign(params);
            /*
            * appId=appId&creatorId=test&expire=1702868725170&signature=BBAA15DDCBF9F8366308331015B88945A353E487
            * */
            System.out.println("生成的签名串:"+sign);
            // 获取到sign后调用创建接口 https://infi.cn/static/sdk_doc/openapi/board_api/create_board.html
            HashMap<String, Object> emptyParams = new HashMap<String,Object>();
            // 调用接口创建白板，将接口路径拼接上之前生成的sign即可
            String createBoardRs= HttpUtil.post("https://dev.plaso.cn/u3wbs/wbs/nc/createBoard?"+sign,emptyParams);
            /*
            * {"code":0,"reqId":"req60","reqTime":1702868657525,"obj":{"recordId":"657fb6b1e9ebaa0001bfa7c1"}}
            * 其中recordId为白板唯一ID,需要保存起来，后续针对白板的操作都需要用到
            * */
            System.out.println("创建白板接口:"+createBoardRs);

            /*
            *
            * #########################################以下为生成连接白板所需用的querySign#############################################
            *
            * */
            // 初始化签名对象
            InfiBoardQuerySignDemo infiboardQuerySignDemo = new InfiBoardQuerySignDemo("appId","signkey");
            HashMap<String, Object> infiBoardQuerySignParams = new HashMap<String,Object>();
            // 此处构造接口文档所需的参数
            // 除appId expire signature 这些参数由sign函数自行注入
            // 以下几个参数皆为必填字段
            infiBoardQuerySignParams.put("recordId","657fb6b1e9ebaa0001bfa7c1"); // 白板唯一ID,即上述接口创建获取到的recordId
            infiBoardQuerySignParams.put("ownerLoginName","test"); // 白板创建者的唯一用户名，与创建白板时使用的creatorId一致
            infiBoardQuerySignParams.put("loginName","test"); // 白板唯一用户名，为对接系统中的用户唯一标识，后续相关回调以及白板中确定用户身份都会用到
            infiBoardQuerySignParams.put("userName","test"); // 白板中光标显示的用户实际名称
            infiBoardQuerySignParams.put("userType","editor"); // editor:编辑 owner:编辑 visitor:只可查看
            infiBoardQuerySignParams.put("opDays","180"); // 操作保留天数,默认180即可
            infiBoardQuerySignParams.put("versionDays","180"); // 历史版本保留天数 默认180即可
            infiBoardQuerySignParams.put("crypto","1"); // 是否加密 1:加密,必须填1
            String infiBoardQuerySign = infiboardQuerySignDemo.getInfiBoardQuerySign(infiBoardQuerySignParams);
            // 获取到连接白板的query,后端下发给前端，前端注入到英飞白板WebSDK中即可
            // appId=appId&crypto=1&loginName=test&opDays=180&ownerLoginName=test&recordId=657fb6b1e9ebaa0001bfa7c1&userName=test&userType=editor&validBegin=1703551773&validTime=120&versionDays=180&signature=24B8DDAF36A7CFFCA51109400DCB2FEDC6D608E9
            System.out.println("生成的白板连接签名串:"+infiBoardQuerySign);
        } catch (Exception e) {
            System.out.println(e.getMessage());
        }
    }
}