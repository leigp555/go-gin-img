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
	rdb := global.Redb
	//json验证
	err := c.ShouldBind(&userEmail)
	if err != nil {
		msg := utils.GetValidMsg(err, &userEmail)
		c.JSON(400, gin.H{"code": 400, "errors": map[string]any{"body": []string{msg}}})
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
	err = rdb.Set(ctx, userEmail.Email, randStr, 300*time.Second).Err()
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "errors": map[string]any{"body": []string{"服务器异常请重试"}}})
		return
	}
	//发送验证码
	err = utils.Email.Send([]string{userEmail.Email}, randStr)
	if err != nil {
		fmt.Println("email发送失败")
		c.JSON(500, gin.H{"msg": "邮件发送失败请重试"})
	} else {
		c.JSON(200, gin.H{"msg": "邮件已发送请注意查收"})
	}
}
