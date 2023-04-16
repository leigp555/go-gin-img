package test

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
	"testing"
)

type C struct{}

// configJsonBody json request body.
type configJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

var store = base64Captcha.DefaultMemStore

// Generate  生成验证码
func (C) Generate() (string, string, error) {
	var param = configJsonBody{
		Id:          "",
		CaptchaType: "",
		VerifyValue: "",
		DriverAudio: &base64Captcha.DriverAudio{},
		DriverString: &base64Captcha.DriverString{
			Length:          4,
			Height:          40,
			Width:           100,
			ShowLineOptions: 2,
			NoiseCount:      0,
			Source:          "1234567890",
		},
		DriverChinese: &base64Captcha.DriverChinese{},
		DriverMath: &base64Captcha.DriverMath{
			Height: 40,
			Width:  100,
		},
		DriverDigit: &base64Captcha.DriverDigit{
			Height:   40,
			Width:    100,
			Length:   4,
			MaxSkew:  0.4,
			DotCount: 50,
		},
	}
	var driver base64Captcha.Driver
	//create base64 encoding captcha
	switch param.CaptchaType {
	case "audio":
		driver = param.DriverAudio
	case "string":
		driver = param.DriverString.ConvertFonts()
	case "math":
		driver = param.DriverMath.ConvertFonts()
	case "chinese":
		driver = param.DriverChinese.ConvertFonts()
	default:
		driver = param.DriverDigit
	}
	c := base64Captcha.NewCaptcha(driver, store)
	return c.Generate()
}

// Verify  解析验证码
func (C) Verify(id, VerifyValue string) bool {
	return store.Verify(id, VerifyValue, true)
}

func TestCaptcha(t *testing.T) {
	var Captcha = C{}
	captchaId, captchaValue, err := Captcha.Generate()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(captchaId, captchaValue)
	if Captcha.Verify(captchaId, captchaValue) {
		t.Error("验证码错误")
		return
	}
	fmt.Println("验证成功")
}
