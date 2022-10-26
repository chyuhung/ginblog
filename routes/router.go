package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	r1 := r.Group("api/v1")
	{
		// User 模块路由接口
		// 添加用户
		r1.POST("/user/add", v1.AddUser)
		// 所有用户列表
		r1.GET("users", v1.GetUsers)
		// 编辑用户
		r1.PUT("user/:id", v1.EditUser)
		// 删除用户
		r1.DELETE("user/:id", v1.DeleteUser)

		// Login 模块路由接口

		// Category 模块路由接口
		// 添加分类
		r1.POST("/category/add", v1.AddCategory)
		// 所有分类列表
		r1.GET("categories", v1.GetCategory)
		// 编辑分类
		r1.PUT("category/:id", v1.EditCategory)
		// 删除分类
		r1.DELETE("category/:id", v1.DeleteCategory)

		// Article 模块路由接口
		// 添加文章
		r1.POST("/article/add", v1.AddArticle)
		// 所有文章列表
		r1.GET("articles", v1.GetArticle)
		// 编辑文章
		r1.PUT("article/:id", v1.EditArticle)
		// 删除文章
		r1.DELETE("article/:id", v1.DeleteArticle)
		// 查询分类下所有文章
		r1.GET("article/list/:id", v1.GetCategoryArticleInfo)
		// 查询单个文章
		r1.GET("article/:id", v1.GetArticleInfo)

		// 测试
		r1.GET("/test", func(context *gin.Context) {
			context.JSONP(http.StatusOK, gin.H{"msg": "ok"})
		})
	}
	r.Run(utils.HttpPort)
}
