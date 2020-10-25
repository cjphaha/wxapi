package wxapi

type Account struct {
	appID string
	appSecret string
}

func NewAccount(appID,appSecret string) *Account{
	return &Account{
		appID: appID,
		appSecret: appSecret,
	}
}

//创建普通微信客户端
func NewClient(account *Account) *CommonClient{
	return &CommonClient{
		account: account,
		signType: MD5,
	}
}


