package main

import (
	"mygram/configs"
	"mygram/controllers"
	"mygram/database"
	"mygram/repositories"
	"mygram/routers"
	"mygram/services"

	"github.com/gin-gonic/gin"
)

// @title MyGram
// @version 1.0
// @description MyGram is a simple social media for saving photos or making comments on other people's photos
// @contact.name Maulana Dwi Wahyudi
// @contact.email maulana@email.com
// @host localhost:8080
// @BasePath /
func main() {
	configs.GetConfig()
	server := gin.Default()
	db := database.GetDB()

	// User DI
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)
	userRouter := routers.NewUserRouter(server, userController)
	userRouter.Setup()

	// Photo
	photoRepo := repositories.NewPhotoRepository(db)
	photoService := services.NewPhotoService(photoRepo)
	photoController := controllers.NewPhotoController(photoService)
	photoRouter := routers.NewPhotoRouter(server, photoController)
	photoRouter.Setup()

	// Comment
	commentRepo := repositories.NewCommentRepository(db)
	commentService := services.NewCommentRepository(commentRepo)
	commentController := controllers.NewCommentController(commentService)
	commentRouter := routers.NewCommentRouter(server, commentController)
	commentRouter.Setup()

	// Social Media
	socialMediaRepo := repositories.NewSocialMediaRepository(db)
	socialMediaService := services.NewSocialMediaService(socialMediaRepo)
	socialMediaController := controllers.NewSocialMediaController(socialMediaService)
	socialMediaRouter := routers.NewSocialMediaRouter(server, socialMediaController)
	socialMediaRouter.Setup()

	// Swagger
	swaggerRouter := routers.NewSwaggerRouter(server)
	swaggerRouter.Start()

	server.Run(configs.App.AppPort)
}
