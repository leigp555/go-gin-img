package public_api

import (
	"github.com/gin-gonic/gin"
	"img/server/utils"
)

func (PublicApi) SendImgCaptcha(c *gin.Context) {
	res := utils.Res
	id, captcha, err := utils.Captcha.Generate()
	if err != nil {
		res.Fail.Error(c, err, "图形验证码生成失败")
		return
	}
	res.Success.Normal(c, "验证码获取成功", map[string]any{"captchaId": id, "captchaImg": captcha})
}
