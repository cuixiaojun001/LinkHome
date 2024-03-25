package qiniu

import "errors"

var Client *QiniuClient

func Init(config map[string]interface{}) error {
	accessKey, ok := config["access_key"].(string)
	if !ok {
		return errors.New("[config] qiniu miss access_key parameter")
	}
	secretKey, ok := config["secret_key"].(string)
	if !ok {
		return errors.New("[config] qiniu miss secret_key parameter")
	}
	domain, ok := config["domain"].(string)
	if !ok {
		return errors.New("[config] qiniu miss domain parameter")
	}
	bucket, ok := config["bucket"].(string)
	if !ok {
		return errors.New("[config] qiniu miss bucket parameter")
	}
	Client = &QiniuClient{
		AccessKey: accessKey,
		SecretKey: secretKey,
		Domain:    domain,
		Bucket:    bucket,
	}
	return nil
}
