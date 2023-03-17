package user_api

import "github.com/gin-gonic/gin"

func (UserApi) UserLogin(c *gin.Context) {
	c.JSON(200, gin.H{"tom": "login"})
}

func (UserApi) UserRegister(c *gin.Context) {
	c.JSON(200, gin.H{"tom": "register"})
}
