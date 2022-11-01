package routers

import (
	"mygram/controllers"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
)

type photoRouter struct {
	router          *gin.Engine
	photoController controllers.PhotoController
}

func NewPhotoRouter(g *gin.Engine, pc controllers.PhotoController) *photoRouter {
	return &photoRouter{
		router:          g,
		photoController: pc,
	}
}

func (pr *photoRouter) Setup() {
	photo := pr.router.Group("/photos")
	{
		photo.Use(middlewares.Auth())
		photo.GET("/", pr.photoController.GetPhotos)
		photo.POST("/", pr.photoController.CreatePhoto)
		photo.PUT("/:photoid", pr.photoController.UpdatePhotoByID)
		photo.DELETE("/:photoid", pr.photoController.DeletePhotoByID)
	}
}
