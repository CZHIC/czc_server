package gcos

import (
	"bytes"
	"context"
	"czc-server/library/logger"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/tencentyun/cos-go-sdk-v5"
)

var (
	cosUrl       = g.Cfg().Get("cos.cosUrl").(string)
	cosSecretID  = g.Cfg().Get("cos.cosSecretID").(string)
	cosSecretKey = g.Cfg().Get("cos.cosSecretKey").(string)
	cosBucket    = g.Cfg().Get("cos.cosBucket").(string)
	environment  = g.Cfg().Get("cos.environment").(string)
)

// cos存储
func CosStorageUpLoad(data []byte, cosFileType string, fileName string) (string, string, error) {
	u, _ := url.Parse(cosUrl)
	b := &cos.BaseURL{BucketURL: u}
	// 实例化COS客户端
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  cosSecretID,
			SecretKey: cosSecretKey,
		},
	})
	// 对象键（Key）是对象在存储桶中的唯一标识，这里用environment+cosBucket+cosBucketType+fileName进行区分。
	cosKey := environment + "/" + cosBucket + "-" + cosFileType + "/" + fileName
	logger.Info("CosStorageUpLoad  cosKey  ", cosKey)
	// 1.通过字符串上传对象
	buffer := bytes.NewBuffer(data)
	_, err := c.Object.Put(context.Background(), cosKey, buffer, nil)
	if err != nil {
		logger.Errorf("CosStorageUpLoad ->Object.Put() execute failed. err = %v", err)
		return cosUrl, cosKey, err
	}

	return cosUrl, cosKey, nil
}

// cos下载
func CosStorageDownLoad(urlName string) (*cos.Response, error) {
	//arr := strings.Split(urlName, "/")
	//urlName = strings.Join(arr[3:], "/")
	//urlName = environment + "/" + urlName
	u, _ := url.Parse(cosUrl)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  cosSecretID,
			SecretKey: cosSecretKey,
		},
	})
	return c.Object.Get(context.Background(), urlName, nil)
}

// cos预览
func CosStoragePreview(urlName string) (string, error) {
	u, _ := url.Parse(cosUrl)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  cosSecretID,
			SecretKey: cosSecretKey,
		},
	})
	previewUrl, err := c.Object.GetPresignedURL(context.Background(), http.MethodGet, urlName, cosSecretID,
		cosSecretKey, time.Hour, nil)
	if err != nil {
		logger.Errorf("CosStoragePreview ->Object.GetPresignedURL() execute failed. err = %v", err)
		return "", err
	}
	cosUrl := previewUrl.String()

	res := strings.Split(urlName, ".")

	fileType := res[len(res)-1]
	if fileType == "doc" || fileType == "docx" || fileType == "xlsx" {
		cosUrl += "&ci-process=doc-preview&dstType=pdf" //将文档转pdf,支持预览
	}
	return strings.ReplaceAll(cosUrl, "%2F", "/"), err
}

// 获取文件cosUrl并设置超时时间
func GetFileUrlByCosKey(urlName string, timeOut time.Duration) (string, error) {
	u, _ := url.Parse(cosUrl)
	b := &cos.BaseURL{BucketURL: u}
	// Get preSigned
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  cosSecretID,
			SecretKey: cosSecretKey,
		},
	})
	previewUrl, err := c.Object.GetPresignedURL(context.Background(), http.MethodGet, urlName, cosSecretID,
		cosSecretKey, timeOut, nil)
	if err != nil {
		logger.Errorf("GetFileUrlByCosKey execute failed. err = %v", err)
		return "", err
	}
	return previewUrl.String(), err
}
