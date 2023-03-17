package article_api

import "github.com/gin-gonic/gin"

// Title godoc
// @Summary      List accounts
// @Description  get accounts
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Success      200  {array}   model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts [get]

func (ArticleApi) Title(c *gin.Context) {
	c.JSON(200, gin.H{"title": "pig"})
}

func (ArticleApi) Content(c *gin.Context) {
	c.JSON(200, gin.H{"count": "abc"})
}
