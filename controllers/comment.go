package controllers

import (
	"mygram/helpers"
	"mygram/params"
	"mygram/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController interface {
	CreateComment(ctx *gin.Context)
	GetCommentsByUserID(ctx *gin.Context)
	UpdateCommentByID(ctx *gin.Context)
	DeleteCommentByID(ctx *gin.Context)
}

type commentController struct {
	commentService services.CommentService
}

func NewCommentController(cs services.CommentService) CommentController {
	return &commentController{
		commentService: cs,
	}
}

// @Create Comment
// @tag.name Comment
// @Description API for Adding Comment
// @Accept json
// @Produce json
// @Tags Comment
// @param Body body params.CreateCommentRequest  true "Create Comment"
// @Success 200 {object} params.CreateCommentResponse
// @Router /comments [POST]
func (cc *commentController) CreateComment(ctx *gin.Context) {
	var request params.CreateCommentRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = helpers.ValidateCommentRequest(request.Message)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := int(ctx.Keys["id"].(float64))
	res, err := cc.commentService.CreateComment(userID, request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

// @Get Comment
// @tag.name Comment
// @Description API for Get All Comment
// @Produce json
// @Tags Comment
// @Success 200 {object} params.GetCommentResponse
// @Router /comments [GET]
func (cc *commentController) GetCommentsByUserID(ctx *gin.Context) {
	res, err := cc.commentService.GetCommentsByUserID()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// @Update Comment
// @tag.name Comment
// @Description API for Updating Comment
// @Accept json
// @Produce json
// @Tags Comment
// @param Body body params.UpdateCommentRequest true "Update Comment"
// @Success 200 {object} params.UpdateCommentResponse
// @Router /comments/:commentId [PUT]
func (cc *commentController) UpdateCommentByID(ctx *gin.Context) {
	IDStr := ctx.Param("commentid")
	ID, err := strconv.Atoi(IDStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var request params.UpdateCommentRequest
	err = ctx.ShouldBind(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = helpers.ValidateCommentRequest(request.Message)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := int(ctx.Keys["id"].(float64))
	res, err := cc.commentService.UpdateCommentByID(ID, userID, request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// @Delete Comment
// @tag.name Comment
// @Description API for Deleting Comment
// @Produce json
// @Tags Comment
// @Success 200 {object} params.DeleteCommentResponse
// @Router /comments/:commentId [DELETE]
func (cc *commentController) DeleteCommentByID(ctx *gin.Context) {
	IDStr := ctx.Param("commentid")
	ID, err := strconv.Atoi(IDStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := int(ctx.Keys["id"].(float64))
	res, err := cc.commentService.DeleteCommentByID(ID, userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
