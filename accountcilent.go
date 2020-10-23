package wxapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//获取SeccsionKey
func (c *CommonClient)GetSessiocKey(code string)  UnionIdBody{
	strings :="https://api.weixin.qq.com/sns/jscode2session?appid=" + c.account.appID + "&secret=" + c.account.appSecret +"&js_code=" + code + "&grant_type=authorization_code";
	fmt.Println(strings);
	client := &http.Client{};
	resp, err := client.Get(strings);//get请求
	defer resp.Body.Close();
	body, err := ioutil.ReadAll(resp.Body)//body是返回的数据
	if err != nil {//错误处理
		fmt.Println(err);
	}
	var strrr UnionIdBody; //这里是转化成结构体的
	err = json.Unmarshal(body, &strrr);
	if err != nil {
		fmt.Println("error:",err);
	}
	fmt.Println("session_key:  ",strrr);
	return strrr;
}


func (c *CommonClient) GetDecryptUserInfo (EncryptedInfo EncryptedUserInfo) *WxUserInfo{
	Sessionkey := c.GetSessiocKey(EncryptedInfo.Code)
	return c.Decode(Sessionkey.Session_key,EncryptedInfo.Iv,EncryptedInfo.EncryptedData)
}
