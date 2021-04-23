package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tzr2020/gin_blog/utils"
)

func InitRouter() *gin.Engine {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}

	return r
}
