package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())

	{
		// User 模块路由接口
		// 所有用户列表
		auth.GET("users", v1.GetUsers)
		// 编辑用户
		auth.PUT("user/:id", v1.EditUser)
		// 删除用户
		auth.DELETE("user/:id", v1.DeleteUser)

		// Category 模块路由接口
		// 添加分类
		auth.POST("/category/add", v1.AddCategory)
		// 编辑分类
		auth.PUT("category/:id", v1.EditCategory)
		// 删除分类
		auth.DELETE("category/:id", v1.DeleteCategory)

		// Article 模块路由接口
		// 添加文章
		auth.POST("/article/add", v1.AddArticle)
		// 编辑文章
		auth.PUT("article/:id", v1.EditArticle)
		// 删除文章
		auth.DELETE("article/:id", v1.DeleteArticle)

		// 上传文件
		auth.POST("/upload", v1.Upload)
	}

	router := r.Group("api/v1")
	{
		// User 模块路由接口
		// 添加用户
		router.POST("/user/add", v1.AddUser)

		// Category 模块路由接口
		// 所有分类列表
		router.GET("categories", v1.GetCategory)

		// Article 模块路由接口
		// 所有文章列表
		router.GET("articles", v1.GetArticle)
		// 查询分类下所有文章
		router.GET("article/list/:id", v1.GetCategoryArticleInfo)
		// 查询单个文章
		router.GET("article/:id", v1.GetArticleInfo)

		// Login 模块路由接口
		router.POST("/login", v1.Login)
	}

	r.Run(utils.HttpPort)
}
