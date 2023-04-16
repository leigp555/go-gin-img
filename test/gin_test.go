package test

import (
	"github.com/gin-gonic/gin"
	"log"
	"testing"
)

func TestGin(t *testing.T) {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello gin")
	})
	err := r.Run(":8000")
	if err != nil {
		log.Fatalln(err)
	}
}
