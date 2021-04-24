package routes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/tzr2020/gin_blog/api/v1"
	"github.com/tzr2020/gin_blog/utils"
)

func InitRouter() *gin.Engine {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	rgv1 := r.Group("api/v1")
	{
		// 用户模块的路由接口
		rgv1.POST("user/add", v1.AddUser)
		rgv1.GET("users", v1.GetUsers)
		rgv1.PUT("user/:id", v1.EditUser)
		rgv1.DELETE("user/:id", v1.DeleteUser)

		// 文章模块的路由接口
		rgv1.POST("article/add", v1.AddArticle)
		rgv1.GET("articles", v1.GetArticles)
		rgv1.PUT("article/:id", v1.EditArticle)
		rgv1.DELETE("article/:id", v1.DeleteArticle)

		// 分类模块的路由接口
		rgv1.POST("category/add", v1.AddCategory)
		rgv1.GET("categories", v1.GetCategories)
		rgv1.PUT("category/:id", v1.EditCategory)
		rgv1.DELETE("category/:id", v1.DeleteCategory)
	}

	return r
}
