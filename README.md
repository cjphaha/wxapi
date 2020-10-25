# wxapi

这是用golang编写的一个用来处理微信接口的一个通用包，包含了微信小程序登陆解密信息、创建支付订单、微信网页登陆等接口，目前正在持续更新中。

##  安装

```bash
$ go get github.com/cjphaha/wxapi
```

## go modules

```cgo
require github.com/cjphaha/wxapi
```
## 常量

|    常量名     |                      值                      |         功能         |
| :-----------: | :------------------------------------------: | :------------------: |
|     Sign      |                     sign                     |       标记签名       |
| SessionkeyUrl | https://api.weixin.qq.com/sns/jscode2session | 获取sessionkey的地址 |



## 常用结构体



> 开发者账号结构体

```go
type Account struct {
	appID string
	appSecret string
}
```

> 用session_key获取的

> 微信内网页登陆授权码

```go
type AccessToken struct {
    Access_token  string
    Expires_in    int
    Refresh_token string
    Openid        string
    Scope         string
}
```

## 接口说明

|       函数名       |                功能                |            参数             |    返回值     | 参数说明                         |
| :----------------: | :--------------------------------: | :-------------------------: | :-----------: | -------------------------------- |
|     NewAccount     |          初始化开发者账号          |  (appID,appSecret string)   |   *Account    | 开发者的appid和appserect         |
|     NewClient      |      创建普通微信开发者客户端      |     (account *Account)      | *CommonClient | 账户信息，包括appid和appserect   |
|   GetSessionKey    |           获取sessionkey           | (code ,appid,secret string) |  UnionIdBody  | code是用户在微信小程序中获取的   |
| GetDecryptUserInfo |           获取解密的信息           |  微信小程序提供的加密信息   | 解密后的信息  | 参数由微信小程序提供             |
|  GetWebSessionKey  | 获取在微信网页内登陆时的sessionkey |            code             |  AccessToken  | 微信网页登陆时有微信的js-sdk获取 |

## 样例

#### 微信小程序登陆

在gin框架中定义如下接口

```go
func Register(c *gin.Context){
	var Temp wxapi.EncryptedUserInfo
	c.Bind(&Temp)
	account := wxapi.NewAccount("开发者AppId,"开发者Appsecret")
	client := wxapi.NewClient(account)
	data := client.GetDecryptUserInfo(Temp)
	c.JSON(http.StatusOK,gin.H{
		"DecryptedData":data,
	})
}
```

发送POST请求，请求体为：

```json
{
    "signature":"sadshdiu27ey3w79rfeiusdhaqw21e38red98",
    "encryptedData": "aLHwTFyAnVVzjl1fDNCDoO88WO25ViWAzgRM06bX0Vgq0CQtHHXp2gGdWTUIo1G4y/+4T0U9U9zCwROQKLmCSK4nePShimlsvFpnj0d30YTe/+qeweqwewsdqsdxasxasxadw/dqwdasdasdasdawdasdaasdas/sadr00p5js2IIqwqBXcpux5JgD9O/Tc1teKIJGqe8JkwMUbPsuNVR130n0w+JE8u4QZpiZ4QHUI0zqRLRzdoMnAZFZLkxw2C6BPP+Uj+tqlmSMTlILh/LR0R5sIrABhCPsKmMeeEsVRATQe6XHpkgN7ziPyDJ5ulVZEnu6RYEzYNq300dxSGae1mT/sulkEFcx7bjPs3KiCR1JTcU3FLxNis6vWKsP3OXc4LTWl5ekoDxajwpQ9tGJFscKEp/UPrQ=",
    "iv": "6LqweRFb8UKX/sXT5/C0sQ==",
    "code": "0qweqweqwV1Ia17JBT821RrGi"
}
```

接口返回的数据为：

```json
{
    "DecryptedData": {
        "id": 0,
        "openId": "oKJKf4udasdaqqweqweaxAiGdS3S2Y",
        "unionId": "o_79T5oh0a2casddYlgQsqIKWn79Y",
        "nickName": "王小二",
        "gender": 1,
        "city": "Nanjing",
        "province": "Jiangsu",
        "country": "China",
        "avatarUrl": "https://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTJB3WhdjBiaQ2t82y3CSFwoeIU8vuWUfiblrHwSpqDlicnLVJWSfzyq42nE8Ok647uAasdasdasbjlNibvEA/132",
        "language": "zh_CN",
        "watermark": {
            "timestamp": 16034352582,
            "appid": "wasasdj3599dWsz3b"
        }
    }
}
```

#### 微信网页登陆

```go
func Register(c *gin.Context){
	account := wxapi.NewAccount("开发者AppId,"开发者Appsecret")
	client := wxapi.NewClient(account)
  code := c.Query("code")
	data := client.GetWebUserInfo(code)
	c.JSON(http.StatusOK,gin.H{
		"UserInfo":data,
	})
}
```




