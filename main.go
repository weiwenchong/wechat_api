package main

import (
	"github.com/gin-gonic/gin"
	"wechat_api/handler"
)

func main() {
	r := gin.Default()

	wxGroup := r.Group("/wx")
	{
		wxGroup.GET("", handler.Auth)
		wxGroup.POST("", handler.Post)
	}

	gin.Logger()

	r.Run("0.0.0.0:80")
}
