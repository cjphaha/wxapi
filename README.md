# wxapi

这是用golang编写的一个用来处理微信接口的一个通用包，包含了微信小程序登陆解密信息、创建支付订单、微信网页登陆等接口，目前正在持续更新中。

##  安装

```bash
$ go get github.com/cjphaha/wxapi
```

## go modules

```cgo
require github.com/cjphaha/wxapi v0.0.2
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

## 接口说明

|       函数名       |            功能             |            参数             |    返回值     | 参数说明                       |
| :----------------: | :-------------------------: | :-------------------------: | :-----------: | ------------------------------ |
|     NewAccount     |      初始化开发者账号       |  (appID,appSecret string)   |   *Account    | 开发者的appid和appserect       |
|     NewClient      |  创建普通微信开发者客户端   |     (account *Account)      | *CommonClient | 账户信息，包括appid和appserect |
|       Jiemi        | 解密用户的信息，获取UnionID |                             |               |                                |
|   GetSessionKey    |       获取sessionkey        | (code ,appid,secret string) |  UnionIdBody  | code是用户在微信小程序中获取的 |
| GetDecryptUserInfo |       获取解密的信息        |  微信小程序提供的加密信息   | 解密后的信息  | 参数由微信小程序提供           |

## 样例








