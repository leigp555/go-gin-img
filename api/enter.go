package api

import (
	"img/server/api/image_api"
	"img/server/api/public_api"
	"img/server/api/socket_api"
	"img/server/api/user_api"
)

type apiGroup struct {
	User   user_api.UserApi
	Public public_api.PublicApi
	Img    image_api.ImgApi
	Socket socket_api.SocketApi
}

var Group = new(apiGroup)
