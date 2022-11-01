package routers

import (
	"mygram/controllers"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
)

type socialMediaRouter struct {
	router                *gin.Engine
	socialMediaController controllers.SocialMediaController
}

func NewSocialMediaRouter(r *gin.Engine, smc controllers.SocialMediaController) *socialMediaRouter {
	return &socialMediaRouter{
		router:                r,
		socialMediaController: smc,
	}
}

func (smr *socialMediaRouter) Setup() {
	socialMedia := smr.router.Group("/socialmedias")
	{
		socialMedia.Use(middlewares.Auth())
		socialMedia.POST("/", smr.socialMediaController.CreateSocialMedia)
		socialMedia.GET("/", smr.socialMediaController.GetSocialMedias)
		socialMedia.PUT("/:socialMediaId", smr.socialMediaController.UpdateSocialMediaByID)
		socialMedia.DELETE("/:socialMediaId", smr.socialMediaController.DeleteSocialMediaByID)
	}
}
