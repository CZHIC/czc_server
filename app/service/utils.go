package service

import (
	"bytes"
	"crypto/sha256"
	"czc-server/library/logger"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gogf/gf/frame/g"
)

// 获取token
func getToken() string {
	return g.Cfg().GetString("rio.token")
}

// 获取签名
func GetSn() (int64, string, string) {
	// 秒级时间戳
	timestamp := GetTimestamp() / 1e3
	// sn
	sn := fmt.Sprintf("%d%s%d", timestamp, getToken(), timestamp)
	signature := fmt.Sprintf("%x", sha256.Sum256([]byte(sn)))
	return timestamp, sn, signature
}

// 获取Moa登录签名
func GetMoaLoginSn() string {
	return g.Cfg().GetString("setting.SystemKey")
}

// 获取Moa登录鉴权url
func GetMoaLoginUrl() string {
	return g.Cfg().GetString("setting.MoaLoginUrl")
}

// 获取 httpRequest
func GetReq(url string, method string, data []interface{}) (*http.Request, error) {
	timestamp, _, signature := GetSn()
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		logger.Error(err)
		return nil, err
	}
	//logger.Debug(fmt.Sprint(timestamp))
	//logger.Debug(signature)

	req.Header.Set("timestamp", fmt.Sprint(timestamp))
	req.Header.Set("signature", signature)

	return req, nil
}

// 获取 httpRequest
func PostReq(url string, method string, data interface{}) (*http.Request, error) {
	timestamp, _, signature := GetSn()
	postData, err := json.Marshal(data)
	var jsonStr = []byte(postData) //转换二进制
	buffer := bytes.NewBuffer(jsonStr)
	req, err := http.NewRequest(method, url, buffer)

	if err != nil {
		logger.Error(err)
		return nil, err
	}
	logger.Info("PostReq  url   ", url, "  timestamp    ", fmt.Sprint(timestamp),
		"   signature   ", signature, "   data   ", data)
	req.Header.Set("timestamp", fmt.Sprint(timestamp))
	req.Header.Set("signature", signature)
	req.Header.Set("Content-Type", "application/json")
	logger.Info("req.Header   ", req.Header)
	return req, nil
}

// 执行请求
func RunReq(req *http.Request) (string, error) {
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		logger.Error(err)
		return "", err
	}
	buf := new(strings.Builder)
	_, err = io.Copy(buf, res.Body)

	if err != nil {
		logger.Error(err)
		return "", err
	}
	logger.Debug("RunReq body", buf.String())
	return buf.String(), nil
}

// 串联 GetReq 与 RunReq, 执行完整请求
func Do(url string, method string, body []interface{}, res interface{}) error {
	req, err := GetReq(url, method, body)
	if err != nil {
		return err
	}

	data, err := RunReq(req)

	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(data), res); err != nil {
		return err
	}

	return nil
}

// 获取 httpRequest
func GetMoaLoginReq(url string, method string, data []interface{}) (*http.Request, error) {
	signature := GetMoaLoginSn()
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	req.Header.Set("X-System-Key", signature)
	return req, nil
}

// 获取 httpRequest
func PostMoaLoginReq(url1 string, method string, reqMap map[string]string) (*http.Request, error) {
	signature := GetMoaLoginSn()
	//postData, err := json.Marshal(data)
	dataVal := url.Values{}
	for i, v := range reqMap {
		dataVal.Add(i, v)
	}
	logger.Info(dataVal)
	//var jsonStr = []byte(postData) //转换二进制
	//buffer := bytes.NewBuffer(jsonStr)

	req, err := http.NewRequest(method, url1, strings.NewReader(dataVal.Encode()))

	if err != nil {
		logger.Error(err)
		return nil, err
	}
	req.Header.Set("X-System-Key", signature)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-Protocol-Version", "ITLoginV1.1")
	logger.Info("req.Header   ", req.Header)
	return req, nil
}

func GetTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}
