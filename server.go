package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/go_restful_api/controllers"
	"github.com/ydhnwb/go_restful_api/repositories"
	"github.com/ydhnwb/go_restful_api/services"
)

var (
	bookRepository repositories.BookRepository = repositories.NewBookRepository()
	bookService    services.BookService        = services.NewBookService(bookRepository)
	loginService   services.LoginService       = services.NewLoginService()
	jwtService     services.JWTService         = services.NewJWTService()

	bookController  controllers.BookController  = controllers.NewBookController(bookService)
	loginController controllers.LoginController = controllers.NewLoginController(loginService, jwtService)
)

func main() {
	defer bookRepository.CloseDatabaseConnection()
	server := gin.Default()

	authRoutes := server.Group("api/auth")
	{
		authRoutes.POST("/login", loginController.Login)
	}

	videosRoutes := server.Group("api/books")
	{
		videosRoutes.GET("/", bookController.All)
		videosRoutes.GET("/:id", bookController.FindByID)
		videosRoutes.POST("/", bookController.Insert)
		videosRoutes.PUT("/:id", bookController.Update)
		videosRoutes.DELETE("/:id", bookController.Delete)
	}

	server.Run(":8080")
}
