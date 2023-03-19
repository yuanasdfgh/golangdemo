package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/api/v1/chat"
	"server/middleware"
)

func Router() {

	engine := gin.Default()

	// 开启跨域
	engine.Use(middleware.Cors())

	// 静态资源请求映射
	//engine.Static("/image", global.Config.Upload.SavePath)

	// 404
	engine.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found")
	})

	// 后台管理API
	web := engine.Group("/web")

	{
		web.GET("/test")
		// 用户登录
		//web.GET("/captcha", api.GetWebUser().GetCaptcha)
		//web.POST("/login", api.GetWebUser().UserLogin)

		// 开启JWT认证,以下接口需要认证成功才能访问
		web.Use(middleware.JwtAuth())

	}

	// 微信小程序API
	app := engine.Group("/app")

	{
		//app.POST("/chatgpt", func(c *gin.Context) {
		//
		//	des := c.PostForm("des")
		//
		//	client := gogpt.NewClient("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")
		//	ctx := context.Background()
		//
		//	req := gogpt.CompletionRequest{
		//		Model:     gogpt.GPT3TextDavinci003,
		//		MaxTokens: 2000,
		//		Prompt:    des,
		//	}
		//	fmt.Println(des, "des")
		//	resp, err := client.CreateCompletion(ctx, req)
		//	if err != nil {
		//		return
		//	}
		//	fmt.Println(resp.Choices[0].Text)
		//	response.Success("操作成功", resp, c)
		//})
		app.GET("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "hello word")
		})
		app.GET("/chatSession", chat.GetChatSession)
		app.GET("/conversation", chat.GetConversation)
		app.POST("/getGpt", chat.GetGpt)
		// 用户登录
		//app.POST("/login", api.GetAppUser().UserLogin)
	}

	// 启动、监听端口
	_ = engine.Run(fmt.Sprintf(":%s", "8000"))
}
