package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/go_restful_api/controllers"
	"github.com/ydhnwb/go_restful_api/repositories"
	"github.com/ydhnwb/go_restful_api/services"
)

var (
	videoRepository repositories.VideoRepository = repositories.NewVideoRepository()
	videoService    services.VideoService        = services.NewVideoService(videoRepository)
	loginService    services.LoginService        = services.NewLoginService()
	jwtService      services.JWTService          = services.NewJWTService()

	videoController controllers.VideoController = controllers.NewVideoController(videoService)
	loginController controllers.LoginController = controllers.NewLoginController(loginService, jwtService)
)

func main() {
	defer videoRepository.CloseDatabaseConnection()
	server := gin.Default()

	authRoutes := server.Group("api/auth")
	{
		authRoutes.POST("/login", loginController.Login)
	}

	videosRoutes := server.Group("api/videos")
	{
		videosRoutes.GET("/", videoController.All)
		videosRoutes.POST("/", videoController.Insert)
		videosRoutes.PUT("/:id", videoController.Update)
		videosRoutes.DELETE("/:id", videoController.Delete)
	}

	server.Run(":8080")
}
