package test

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"testing"
)

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	}
}

func TestWebSocket(t *testing.T) {
	r := gin.Default()
	r.Use(cors())
	r.GET("/chat", ws)
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var coons []*websocket.Conn

func ws(c *gin.Context) {
	// upgrade to websocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	coons = append(coons, conn)

	//获取socket内容,并发送给客户端
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(string(message))

		for _, conn := range coons {
			if err := conn.WriteMessage(messageType, message); err != nil {
				log.Println(err)
				break
			}
		}

	}
}
