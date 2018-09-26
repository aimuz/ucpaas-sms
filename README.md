# ucpaas-sms
云之讯短信go版SDK

```go
func main() {
    config := &ucpaassmsclient.Config{
    	SID:"39467b989d087c2d92c6132184a365d8", // 用户的账号唯一标识“Account Sid”，在开发者控制台获取
    	Token:"23f757bad208226ec301e117e40006ed", // 用户密钥“Auth Token”，在开发者控制台获取
    	AppID:"2d92c6132139467b989d087c84a365d8", // 创建应用时系统分配的唯一标示
    }
    config.Send("模板ID","参数","手机号","uid")
}
```
