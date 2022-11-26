package main

import (
	"github.com/gin-gonic/gin"
	"wechat_api/handler"
	"wechat_api/log"
)

func main() {
	r := gin.Default()

	wxGroup := r.Group("/wx")
	{
		wxGroup.GET("", handler.Auth)
		wxGroup.POST("", handler.Post)
	}

	log.Infof("service start")

	r.Run("0.0.0.0:80")
}
