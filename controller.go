package controller

/* assets的资源文件：https://getbootstrap.com/ https://pandao.github.io/editor.md/ */
import (
	"blog/dao"
	"blog/model"
	"fmt"
	"html/template"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

func RegisterUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	password2 := c.PostForm("password2")
	if username != "" && password != "" {
		if password == password2 {
			user := model.User{
				Username: username,
				Password: password,
			}

			dao.Mgr.Redirect(&user)
			c.Redirect(301, "/")
		} else {
			c.String(200, "密码不一致")
		}
	} else {
		c.String(200, "名称与密码不能为空")
	}
}
func GORegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
func GoLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username)
	u := dao.Mgr.Login(username)

	if u.Username == "" {
		c.HTML(200, "login.html", "用户名不存在！")
		fmt.Println("用户名不存在！")
	} else {
		if u.Password != password {
			fmt.Println("密码错误")
			c.HTML(200, "login.html", "密码错误")
		} else {
			fmt.Println("登录成功")
			c.Redirect(301, "/")
		}
	}

}

//操作博客

// 博客列表
func GetPostIndex(c *gin.Context) {
	posts := dao.Mgr.GetAllPost()
	c.HTML(200, "postIndex.html", posts)
}

// 添加博客
func AddPost(c *gin.Context) {
	title := c.PostForm("title")
	tag := c.PostForm("tag")
	content := c.PostForm("content")

	post := model.Post{
		Title:   title,
		Tag:     tag,
		Content: content,
	}

	dao.Mgr.AddPost(&post)

	c.Redirect(302, "/post_index")
}

// 跳转到添加博客
func GoAddPost(c *gin.Context) {
	c.HTML(200, "post.html", nil)
}

// 博客详细控制器
func PostDetail(c *gin.Context) {
	s := c.Query("pid")
	pid, _ := strconv.Atoi(s)
	p := dao.Mgr.GetPost(pid)

	content := blackfriday.Run([]byte(p.Content))
	c.HTML(200, "detail.html", gin.H{
		"Title":   p.Title,
		"Content": template.HTML(content),
	})
}
