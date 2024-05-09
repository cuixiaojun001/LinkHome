package qiniu

import (
	"bytes"
	"context"
	"fmt"
	"github.com/cuixiaojun001/linkhome/library/utils"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
	"time"
)

type QiniuClient struct {
	AccessKey string
	SecretKey string
	Domain    string
	Bucket    string
}

func (c *QiniuClient) MakePrivateURL(key string) string {
	mac := auth.New(c.AccessKey, c.SecretKey)
	// 生成一个私有空间的下载链接
	deadline := time.Now().Add(time.Second * 3600).Unix() //1小时有效期
	privateAccessURL := storage.MakePrivateURL(mac, c.Domain, key, deadline)
	return privateAccessURL
}

func (c *QiniuClient) UploadFile(data []byte) (key, url string) {
	putPolicy := storage.PutPolicy{
		Scope: c.Bucket,
	}
	mac := auth.New(c.AccessKey, c.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuabei
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	dataLen := int64(len(data))
	err := formUploader.Put(context.Background(), &ret, upToken, utils.GenerateRandomString(28), bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	key = ret.Key
	url = c.Domain + "/" + key
	return ret.Key, url
}
