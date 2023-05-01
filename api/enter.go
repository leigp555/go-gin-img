package api

import (
	"img/server/api/image_api"
	"img/server/api/public_api"
	"img/server/api/socket_api"
	"img/server/api/user_api"
)

type apiGroup struct {
	UserApi   user_api.UserApi
	PublicApi public_api.PublicApi
	ImgApi    image_api.ImgApi
	SocketApi socket_api.SocketApi
}

var GroupApi = new(apiGroup)
