package public_api

import (
	"github.com/gin-gonic/gin"
	"img/server/utils"
)

func (PublicApi) SendImgCaptcha(c *gin.Context) {
	id, captcha, err := utils.Captcha.Generate()
	if err != nil {
		utils.Res.FailWidthRecord(c, 500, "验证码生成失败", struct{}{}, err, "验证码生成失败")
		return
	}
	utils.Res.Success(c, "验证码发送成功", map[string]any{"captchaId": id, "captchaImg": captcha})
}
