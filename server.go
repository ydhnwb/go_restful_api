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
	videosRoutes := server.Group("/api/videos")
	{
		videosRoutes.GET("/", videoController.FindAll)
		videosRoutes.POST("/", videoController.Save)
	}

	server.Run(":8080")
}
