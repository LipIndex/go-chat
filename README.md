### 宗本聊天消息协议交互说明

* 格式 `json`,以json字符串交互
* 示例
```json
{"action":1,uid:123456,"payload":{"msg":"xxxxxx"}}
```
* 参数说明
  * action：协议号
  * uid:用户id
  * payload：消息体


#### 协议定义
1. 登录
    * action:1

2. 发消息
    * action:2
    * payload
    ``` json
    {
      "msg":"xxxxx"
   }
   ```