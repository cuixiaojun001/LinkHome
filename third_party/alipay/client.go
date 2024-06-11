package alipay

import (
	"github.com/smartwalle/alipay/v3"
)

type AliPayClient struct {
	client      *alipay.Client
	callbackURL string
}

func New(client *alipay.Client, callbackURL string) *AliPayClient {
	return &AliPayClient{
		client:      client,
		callbackURL: callbackURL,
	}
}

func (c *AliPayClient) GetClient() *alipay.Client {
	return c.client
}

func (c *AliPayClient) CallBackURL() string {
	return c.callbackURL
}
