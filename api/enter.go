package api

import "img/server/api/user_api"

type apiGroup struct {
	UserApi user_api.UserApi
}

var GroupApi = new(apiGroup)
