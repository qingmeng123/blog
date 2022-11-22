package v1

import (
	middleware2 "duryun-blog/api/middleware"
	"duryun-blog/config"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("admin", "web/admin/dist/index.html")
	p.AddFromFiles("front", "web/front/dist/index.html")
	return p
}

func InitRouter() {
	gin.SetMode(config.AppMode)
	r := gin.New()
	// 设置信任网络 []string
	// nil 为不计算，避免性能消耗，上线应当设置
	_ = r.SetTrustedProxies(nil)

	r.HTMLRender = createMyRender()
	r.Use(middleware2.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware2.Cors())

	r.Static("/static", "./web/front/dist/static")
	r.Static("/admin", "./web/admin/dist")
	r.StaticFile("/favicon.ico", "/web/front/dist/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})

	/*
		后台管理路由接口
	*/
	auth := r.Group("api/v1")
	auth.Use(middleware2.JwtToken())
	{
		// 用户模块的路由接口
		auth.GET("admin/users", GetUsers)
		auth.PUT("user/:id", EditUser)
		auth.DELETE("user/:id", DeleteUser)
		//修改密码
		auth.PUT("admin/changepw/:id", ChangeUserPassword)
		// 分类模块的路由接口
		auth.GET("admin/category", GetCate)
		auth.POST("category/add", AddCategory)
		auth.PUT("category/:id", EditCate)
		auth.DELETE("category/:id", DeleteCate)
		// 文章模块的路由接口
		auth.GET("admin/article/info/:id", GetArtInfo)
		auth.GET("admin/article", GetArtList)
		auth.POST("article/add", AddArticle)
		auth.PUT("article/:id", EditArt)
		auth.DELETE("article/:id", DeleteArt)
		// 上传文件
		auth.POST("upload", UpLoad)
		// 更新个人设置
		auth.GET("admin/profile/:id", GetProfile)
		auth.PUT("profile/:id", UpdateProfile)
		// 评论模块
		auth.GET("comment/list", GetCommentList)
		auth.DELETE("delcomment/:id", DeleteComment)
		auth.PUT("checkcomment/:id", CheckComment)
		auth.PUT("uncheckcomment/:id", UncheckComment)
	}

	/*
		前端展示页面接口
	*/
	router := r.Group("api/v1")
	{
		// 用户信息模块
		router.POST("user/add", AddUser)
		router.GET("user/:id", GetUserInfo)
		router.GET("users", GetUsers)

		// 文章分类信息模块
		router.GET("category", GetCate)
		router.GET("category/:id", GetCateInfo)

		// 文章模块
		router.GET("article", GetArtList)
		router.GET("article/list/:id", GetCateArt)
		router.GET("article/info/:id", GetArtInfo)

		// 登录控制模块
		router.POST("login", Login)
		router.POST("loginfront", LoginFront)

		// 获取个人设置信息
		router.GET("profile/:id", GetProfile)

		// 评论模块
		router.POST("addcomment", AddComment)
		router.GET("comment/info/:id", GetComment)
		router.GET("commentfront/:id", GetCommentListFront)
		router.GET("commentcount/:id", GetCommentCount)
	}

	_ = r.Run(config.HttpPort)

}
