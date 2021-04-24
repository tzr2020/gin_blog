package routes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/tzr2020/gin_blog/api/v1"
	"github.com/tzr2020/gin_blog/middleware"
	"github.com/tzr2020/gin_blog/utils"
)

func InitRouter() *gin.Engine {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())

	rgv1 := r.Group("api/v1")
	{
		// 用户模块的路由接口
		rgv1.POST("user/add", v1.AddUser)
		rgv1.GET("users", v1.GetUsers)
		rgv1.POST("login", v1.Login)

		// 文章模块的路由接口
		rgv1.GET("articles", v1.GetArticles)
		rgv1.GET("articles/category/:id", v1.GetArticlesByCategory)
		rgv1.GET("article/info/:id", v1.GetArticleInfo)

		// 分类模块的路由接口
		rgv1.GET("categories", v1.GetCategories)

	}

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtMiddleware())
	{
		// 用户模块的路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)

		// 文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)

		// 分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)
	}

	return r
}
