package sms

import (
	"errors"
	"github.com/cloopen/go-sms-sdk/cloopen"
)

var Client *SmsClient

func Init(config map[string]interface{}) error {
	account, ok := config["account"].(string)
	if !ok {
		return errors.New("[config] sms miss account parameter")
	}
	token, ok := config["token"].(string)
	if !ok {
		return errors.New("[config] sms miss token parameter")
	}
	cfg := cloopen.DefaultConfig().WithAPIAccount(account).WithAPIToken(token)
	sms := cloopen.NewJsonClient(cfg).SMS()
	Client = New(sms)
	return nil
}
