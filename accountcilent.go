package wxapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
//获取微信网页登陆时需要用到的sessionkey
func (c *CommonClient)GetWebSessionKey(code string) AccessToken{
	string := "https://api.weixin.qq.com/sns/oauth2/access_token?appid="+ c.account.appID + "&secret=" + c.account.appSecret + "&code=" + code + "&grant_type=authorization_code"
	fmt.Println(string)
	cilent := &http.Client{};
	resp,err := cilent.Get(string)
	defer resp.Body.Close();
	body, err := ioutil.ReadAll(resp.Body)//body是返回的数据
	if err != nil {//错误处理
		fmt.Println(err);
	}
	var strrr AccessToken; //这里是转化成结构体的
	err = json.Unmarshal(body, &strrr);
	if err != nil {
		fmt.Println("error:",err);
	}
	return strrr
}

func (c *CommonClient)GetWebUserInfo(code string){
	access := c.GetWebSessionKey(code)
	strings :=  "https://api.weixin.qq.com/sns/userinfo?access_token=" + access.Access_token + "&openid=" + access.Openid + "&lang=zh_CN"
	cilent1 := &http.Client{};
	resp1,err := cilent1.Get(strings)
	defer resp1.Body.Close();
	body1, err := ioutil.ReadAll(resp1.Body)//body是返回的数据
	if err != nil {//错误处理
		fmt.Println(err);
	}
	var userinfo WebUserInfo;
	err = json.Unmarshal(body1,&userinfo);
}

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
