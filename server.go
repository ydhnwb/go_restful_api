package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/go_restful_api/config"
	"github.com/ydhnwb/go_restful_api/controllers"
	"github.com/ydhnwb/go_restful_api/repositories"
	"github.com/ydhnwb/go_restful_api/services"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                    = config.SetupDatabaseConnection()
	bookRepository repositories.BookRepository = repositories.NewBookRepository(db)
	userRepository repositories.UserRepository = repositories.NewUserRepository(db)
	bookService    services.BookService        = services.NewBookService(bookRepository)
	loginService   services.LoginService       = services.NewLoginService()
	jwtService     services.JWTService         = services.NewJWTService()

	bookController  controllers.BookController  = controllers.NewBookController(bookService)
	loginController controllers.LoginController = controllers.NewLoginController(loginService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
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
