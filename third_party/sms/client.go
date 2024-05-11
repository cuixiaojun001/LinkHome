package sms

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/cloopen/go-sms-sdk/cloopen"
	"github.com/cuixiaojun001/LinkHome/common/cache"
)

type SmsClient struct {
	cache  cache.Cache
	client *cloopen.SMS
}

func New(sms *cloopen.SMS) *SmsClient {
	return &SmsClient{
		cache:  cache.New("sms"),
		client: sms,
	}
}

func (c *SmsClient) SendSmsCode(mobile string, errCh chan<- error) {
	defer close(errCh)
	// cache 处理
	info := mobileSMSCode(mobile)

	var result string
	if exist, _ := c.cache.Get(context.TODO(), info.Key, &result); exist {
		// 5分钟中内重复发送短信验证码不处理
		return
	}
	// 生成验证码
	smsCode := generateSmsCode()
	err := c.cache.SetEX(context.TODO(), info.Key, smsCode, info.Timeout)
	if err != nil {
		errCh <- err
		return
	}
	input := &cloopen.SendRequest{
		// 应用的APPID
		AppId: "2c94811c8cd4da0a018e41aeb0963aba",
		// 手机号码
		To: "19912419238",
		// 模版ID
		TemplateId: "1",
		// 模版变量内容 非必填
		Datas: []string{smsCode, "5"}, // 5分钟有效
	}
	/** 下发
	 * 返回发送结果和发送成功消息ID
	 * 发送成功示例:
	 *   {"statusCode":"000000","templateSMS":{"dateCreated":"20130201155306","smsMessageSid":"ff8080813c373cab013c94b0f0512345"}}
	 * 发送失败示例：
	 *   {"statusCode": "172001", "statusMsg": "网络错误"}
	 */
	resp, err := c.client.Send(input)
	if err != nil {
		log.Println(err)
		errCh <- err
		return
	}
	log.Printf("Response MsgId: %s \n", resp.TemplateSMS.SmsMessageSid)
	errCh <- err
}

// generateSmsCode 生成随机四位数短信验证码
func generateSmsCode() string {
	rand.Seed(time.Now().UnixNano())
	smsCode := rand.Intn(9000) + 1000
	smsCodeStr := strconv.Itoa(smsCode)
	runes := []rune(smsCodeStr)
	rand.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})
	return string(runes)
}

func (c *SmsClient) GetSmsCode(mobile string) string {
	var smsCode string
	if exist, _ := c.cache.Get(context.TODO(), mobileSMSCode(mobile).Key, &smsCode); exist {
		return smsCode
	}
	return ""
}
