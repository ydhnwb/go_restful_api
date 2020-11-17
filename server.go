package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/go_restful_api/controllers"
	"github.com/ydhnwb/go_restful_api/services"
)

var (
	videoService    services.VideoService       = services.NewVideoService()
	videoController controllers.VideoController = controllers.NewVideoController(videoService)
)

func main() {
	server := gin.Default()
	server.GET("/videos", videoController.FindAll)
	server.POST("/videos", videoController.Save)
	server.Run(":8080")
}
