package routers

import (
	"mygram/controllers"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
)

type commentRouter struct {
	router            *gin.Engine
	commentController controllers.CommentController
}

func NewCommentRouter(r *gin.Engine, cc controllers.CommentController) *commentRouter {
	return &commentRouter{
		router:            r,
		commentController: cc,
	}
}

func (cr *commentRouter) Setup() {
	comment := cr.router.Group("/comments")
	{
		comment.Use(middlewares.Auth())
		comment.POST("/", cr.commentController.CreateComment)
		comment.GET("/", cr.commentController.GetCommentsByUserID)
		comment.PUT("/:commentid", cr.commentController.UpdateCommentByID)
		comment.DELETE("/:commentid", cr.commentController.DeleteCommentByID)
	}
}
