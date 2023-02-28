package router

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

/* assets的资源文件：https://getbootstrap.com/ https://pandao.github.io/editor.md/ */
func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")

	e.GET("/login", controller.GoLogin)
	e.POST("/login", controller.Login)
	e.POST("/register", controller.RegisterUser)
	e.GET("/register", controller.GORegister)
	e.GET("/", controller.Index)

	//操作博客

	//博客列表
	e.GET("/post_index", controller.GetPostIndex)
	//添加博客
	e.POST("/post", controller.AddPost)
	//跳转到添加博客
	e.GET("/post", controller.GoAddPost)
	//博客详细控制器
	e.GET("/post_detail", controller.PostDetail)
	e.Run(":8888")
}
