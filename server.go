package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/go_restful_api/controllers"
	"github.com/ydhnwb/go_restful_api/services"
)

var (
	videoService services.VideoService = services.NewVideoService()
	loginService services.LoginService = services.NewLoginService()
	jwtService   services.JWTService   = services.NewJWTService()

	videoController controllers.VideoController = controllers.NewVideoController(videoService)
	loginController controllers.LoginController = controllers.NewLoginController(loginService, jwtService)
)

func main() {
	server := gin.Default()

	authRoutes := server.Group("api/auth")
	{
		authRoutes.POST("/login", loginController.Login)
	}

	videosRoutes := server.Group("api/videos")
	{
		videosRoutes.GET("/", videoController.FindAll)
		videosRoutes.POST("/", videoController.Save)
	}

	server.Run(":8080")
}
