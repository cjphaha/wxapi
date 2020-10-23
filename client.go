package wxapi

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

const bodyType = "application/xml; charset=utf-8"

type PayClient struct {
	account              *PayAccount // 支付账号
	signType             string   // 签名类型
	httpConnectTimeoutMs int      // 连接超时时间
	httpReadTimeoutMs    int      // 读取超时时间
}

type CommonClient struct {
	account *Account
	signType string
}

// 创建微信支付客户端
func NewPayClient(account *PayAccount) *PayClient {
	return &PayClient{
		account:              account,
		signType:             MD5,
		httpConnectTimeoutMs: 2000,
		httpReadTimeoutMs:    1000,
	}
}

//创建普通微信客户端
func NewClient(account *Account) *CommonClient{
	return &CommonClient{
		account: account,
		signType: MD5,
	}
}

func (c *PayClient) SetHttpConnectTimeoutMs(ms int) {
	c.httpConnectTimeoutMs = ms
}

func (c *PayClient) SetHttpReadTimeoutMs(ms int) {
	c.httpReadTimeoutMs = ms
}

func (c *PayClient) SetSignType(signType string) {
	c.signType = signType
}

func (c *PayClient) SetAccount(account *PayAccount) {
	c.account = account
}

// 向 params 中添加 appid、mch_id、nonce_str、sign_type、sign
func (c *PayClient) fillRequestData(params Params) Params {
	params["appid"] = c.account.appID
	params["mch_id"] = c.account.mchID
	params["nonce_str"] = nonceStr()
	params["sign_type"] = c.signType
	params["sign"] = c.Sign(params)
	return params
}

// https no cert post
func (c *PayClient) postWithoutCert(url string, params Params) (string, error) {
	h := &http.Client{}
	p := c.fillRequestData(params)
	response, err := h.Post(url, bodyType, strings.NewReader(MapToXml(p)))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// 验证签名
func (c *PayClient) ValidSign(params Params) bool {
	if !params.ContainsKey(Sign) {
		return false
	}
	return params.GetString(Sign) == c.Sign(params)
}

// 签名
func (c *PayClient) Sign(params Params) string {
	// 创建切片
	var keys = make([]string, 0, len(params))
	// 遍历签名参数
	for k := range params {
		if k != "sign" { // 排除sign字段
			keys = append(keys, k)
		}
	}
	// 由于切片的元素顺序是不固定，所以这里强制给切片元素加个顺序
	sort.Strings(keys)
	//创建字符缓冲
	var buf bytes.Buffer
	for _, k := range keys {
		if len(params.GetString(k)) > 0 {
			buf.WriteString(k)
			buf.WriteString(`=`)
			buf.WriteString(params.GetString(k))
			buf.WriteString(`&`)
		}
	}
	// 加入apiKey作加密密钥
	buf.WriteString(`key=`)
	buf.WriteString(c.account.apiKey)
	var (
		dataMd5    [16]byte
		dataSha256 []byte
		str        string
	)
	switch c.signType {
	case MD5:
		dataMd5 = md5.Sum(buf.Bytes())
		str = hex.EncodeToString(dataMd5[:]) //需转换成切片
	case HMACSHA256:
		h := hmac.New(sha256.New, []byte(c.account.apiKey))
		h.Write(buf.Bytes())
		dataSha256 = h.Sum(nil)
		str = hex.EncodeToString(dataSha256[:])
	}
	return strings.ToUpper(str)
}

// 处理 HTTPS API返回数据，转换成Map对象。return_code为SUCCESS时，验证签名。
func (c *PayClient) processResponseXml(xmlStr string) (Params, error) {
	var returnCode string
	params := XmlToMap(xmlStr)
	if params.ContainsKey("return_code") {
		returnCode = params.GetString("return_code")
	} else {
		return nil, errors.New("no return_code in XML")
	}
	if returnCode == Fail {
		return params, nil
	} else if returnCode == Success {
		if c.ValidSign(params) {
			return params, nil
		} else {
			return nil, errors.New("invalid sign value in XML")
		}
	} else {
		return nil, errors.New("return_code value is invalid in XML")
	}
}

// 统一下单
func (c *PayClient) UnifiedOrder(params Params) (Params, error) {
	var url string
	if c.account.isSandbox {
		url = SandboxUnifiedOrderUrl
	} else {
		url = UnifiedOrderUrl
	}
	xmlStr, err := c.postWithoutCert(url, params)
	if err != nil {
		return nil, err
	}
	return c.processResponseXml(xmlStr)
}

// 订单查询
func (c *PayClient) OrderQuery(params Params) (Params, error) {
	var url string
	if c.account.isSandbox {
		url = SandboxOrderQueryUrl
	} else {
		url = OrderQueryUrl
	}
	xmlStr, err := c.postWithoutCert(url, params)
	if err != nil {
		return nil, err
	}
	return c.processResponseXml(xmlStr)
}
