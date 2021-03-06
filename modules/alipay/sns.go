package alipay

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/go-baa/log"
)

// SnsLoginInfo 小程序登录后 信息
type SnsLoginInfo struct {
	OauthTokenTesponse oauthTokenResponse `json:"alipay_system_oauth_token_response"`
	ErrorResponse      errorResponse      `json:"error_response"`
}

// SnsLogin .
func SnsLogin(appID, pk, authCode string) (*SnsLoginInfo, error) {
	biz := bizOauthTokenContent{GrantType: "authorization_code", Code: authCode}
	bizCode, _ := json.Marshal(&biz)
	params := url.Values{}
	params.Add("app_id", appID)
	params.Add("biz_content", string(bizCode))
	params.Add("charset", "utf-8")
	params.Add("code", authCode)
	params.Add("grant_type", "authorization_code")
	params.Add("method", oauthTokenAPI)
	params.Add("sign_type", "RSA2")
	params.Add("timestamp", time.Now().Format("2006-01-02 15:04:05"))
	params.Add("version", "1.0")
	sign := sign(sortKeys(params), pk)
	params.Add("sign", sign)
	client := &http.Client{
		Timeout: time.Second * time.Duration(APIRequestTimeout),
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	fmt.Println(params)
	resp, err := client.PostForm(alipayGateway+"?"+params.Encode(), params)
	if err != nil {
		log.Errorf("请求支付宝小程序接口 失败：%s", err.Error())
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("支付宝小程序登录 失败：%s", err.Error())
		return nil, err
	}
	ret := new(SnsLoginInfo)
	if err := json.Unmarshal(data, ret); err != nil {
		return nil, err
	}
	if ret.ErrorResponse.Code != "" {
		return nil, fmt.Errorf("支付宝小程序登录 支付宝API响应错误: %s %s", ret.ErrorResponse.Code, ret.ErrorResponse.Msg)
	}
	return ret, nil
}

type phone struct {
	Code   string `json:"code"`
	Msg    string `json:"msg"`
	Mobile string `json:"mobile"`
}

// GetPhone .
func GetPhone(ociphertext, okey string) (string, error) {
	bs, err := AESDecrypt(ociphertext, okey)
	if err != nil {
		return "", err
	}
	info := phone{}
	err = json.Unmarshal(bs, &info)
	if err != nil {
		return "", err
	}
	if info.Code != "10000" {
		return "", fmt.Errorf("异常的手机号 %s %s", info.Code, info.Msg)
	}
	return info.Mobile, nil
}
