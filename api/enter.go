package api

import (
	"img/server/api/article_api"
	"img/server/api/image_api"
	"img/server/api/user_api"
)

type apiGroup struct {
	UserApi    user_api.UserApi
	ArticleApi article_api.ArticleApi
	ImgApi     image_api.ImgApi
}

var GroupApi = new(apiGroup)
