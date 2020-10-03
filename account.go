package wxapi

import (
	"io/ioutil"
	"log"
)

type PayAccount struct {
	appID     string
	mchID     string
	apiKey    string
	certData  []byte
	isSandbox bool
}

// 创建微信支付账号
func NewPayAccount(appID string, mchID string, apiKey string, isSanbox bool) *PayAccount {
	return &PayAccount{
		appID:     appID,
		mchID:     mchID,
		apiKey:    apiKey,
		isSandbox: isSanbox,
	}
}

// 设置证书
func (a *PayAccount) SetCertData(certPath string) {
	certData, err := ioutil.ReadFile(certPath)
	if err != nil {
		log.Println("读取证书失败")
		return
	}
	a.certData = certData
}

type CommonAccount struct {
	appID string
	appSecret string
}

func NewCommonAccount(appID,appSecret string) *CommonAccount{
	return &CommonAccount{
		appID: appID,
		appSecret: appSecret,
	}
}


