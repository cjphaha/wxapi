package wxapi

//用户发送的请求
type UserInfo struct {
	NickName string
	Gender int
	City string
	Province string
	Country string
	AvatarUrl string
	Language string
	Code string
	Iv string
	EncryptedData string
}

type WxUserInfo struct {
	Id int64 `json:"id"`
	OpenID    string `json:"openId"`
	UnionID   string `json:"unionId"`
	NickName  string `json:"nickName"`
	Gender    int    `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarURL string `json:"avatarUrl"`
	Language  string `json:"language"`
	Watermark struct {
		Timestamp int64  `json:"timestamp"`
		AppID     string `json:"appid"`
	} `json:"watermark"`
}

type WXUserDataCrypt struct {
	appID, sessionKey string
}

type UnionIdBody struct {
	Session_key string
	Openid     string
}