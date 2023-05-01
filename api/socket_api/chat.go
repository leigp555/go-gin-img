package socket_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"img/server/global"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
	HandshakeTimeout:  10 * time.Second,
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
		fmt.Println(reason)
	},
}

var coons []*websocket.Conn

func (SocketApi) Chat(c *gin.Context) {
	// upgrade to websocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			global.Slog.Fatalln(err)
		}
	}()

	coons = append(coons, conn)

	//获取socket内容,并发送给客户端
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			global.Slog.Error(err)
			break
		}

		for _, conn := range coons {
			if err := conn.WriteMessage(messageType, message); err != nil {
				global.Slog.Error(err)
				break
			}
		}

	}
}
