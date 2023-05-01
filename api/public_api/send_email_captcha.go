package public_api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"img/server/global"
	"img/server/utils"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func (PublicApi) SendEmailCaptcha(c *gin.Context) {
	type Email struct {
		Email string `form:"email" binding:"required,email" msg:"邮箱格式不正确"`
	}
	var userEmail Email
	rdb := global.Rdb
	res := utils.Res
	//json验证
	if err := c.ShouldBind(&userEmail); err != nil {
		msg := utils.GetValidMsg(err, &userEmail)
		res.Fail.Normal(c, 400, msg)
		return
	}
	//生成随机数
	var arr = make([]string, 0)
	for i := 0; i < 6; i++ {
		arr = append(arr, strconv.Itoa(rand.Intn(10)))
	}
	randStr := fmt.Sprintf(strings.Join(arr, ""))
	//存入redis
	var ctx = context.Background()
	if err := rdb.Set(ctx, userEmail.Email, randStr, 300*time.Second).Err(); err != nil {
		res.Fail.Error(c, err, "redis邮箱验证码存储失败")
		return
	}
	//发送验证码
	if err := utils.Email.Send([]string{userEmail.Email}, randStr); err != nil {
		res.Fail.Error(c, err, "邮件发送失败")
		return
	}
	//验证码发送成功的响应
	res.Success.WidthMsg(c, "邮箱验证码发送成功")
}
