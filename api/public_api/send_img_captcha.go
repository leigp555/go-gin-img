package public_api

import (
	"github.com/gin-gonic/gin"
	"img/server/utils"
)

func (PublicApi) SendImgCaptcha(c *gin.Context) {
	id, captcha, err := utils.Captcha.Generate()
	if err != nil {
		utils.DealErr(c, err, "图形验证码生成失败")
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "验证码获取成功", "data": map[string]any{"captchaId": id, "captchaImg": captcha}})
}
