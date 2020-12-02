package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/go_restful_api/config"
	"github.com/ydhnwb/go_restful_api/controllers"
	"github.com/ydhnwb/go_restful_api/middlewares"
	"github.com/ydhnwb/go_restful_api/repositories"
	"github.com/ydhnwb/go_restful_api/services"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                    = config.SetupDatabaseConnection()
	bookRepository repositories.BookRepository = repositories.NewBookRepository(db)
	userRepository repositories.UserRepository = repositories.NewUserRepository(db)
	bookService    services.BookService        = services.NewBookService(bookRepository)
	loginService   services.LoginService       = services.NewLoginService(userRepository)
	jwtService     services.JWTService         = services.NewJWTService()
	userService    services.UserService        = services.NewUserService(userRepository)

	bookController  controllers.BookController  = controllers.NewBookController(bookService)
	loginController controllers.LoginController = controllers.NewLoginController(loginService, jwtService)
	userController  controllers.UserController  = controllers.NewUserController(userService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	server := gin.Default()

	authRoutes := server.Group("api/auth")
	{
		authRoutes.POST("/login", loginController.Login)
		authRoutes.POST("/register", loginController.Register)
	}

	userRoutes := server.Group("api/user", middlewares.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
	}

	bookRoutes := server.Group("api/books", middlewares.AuthorizeJWT(jwtService))
	{
		bookRoutes.GET("/", bookController.All)
		bookRoutes.GET("/:id", bookController.FindByID)
		bookRoutes.POST("/", bookController.Insert)
		bookRoutes.PUT("/:id", bookController.Update)
		bookRoutes.DELETE("/:id", bookController.Delete)
	}

	server.Run(":8080")
}
