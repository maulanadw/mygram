package routers

import (
	"mygram/controllers"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
)

type userRouter struct {
	router         *gin.Engine
	userContorller controllers.UserController
}

func NewUserRouter(r *gin.Engine, uc controllers.UserController) *userRouter {
	return &userRouter{
		router:         r,
		userContorller: uc,
	}
}

func (ur *userRouter) Setup() {
	userRouter := ur.router.Group("/users")
	{
		userRouter.POST("/register", ur.userContorller.Register)
		userRouter.POST("/login", ur.userContorller.Login)
		userRouter.PUT("/:id", middlewares.Auth(), ur.userContorller.UpdateUserByID)
		userRouter.DELETE("/:id", middlewares.Auth(), ur.userContorller.DeleteUserByID)
	}
}
