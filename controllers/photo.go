package controllers

import (
	"mygram/helpers"
	"mygram/params"
	"mygram/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PhotoController interface {
	CreatePhoto(ctx *gin.Context)
	GetPhotos(ctx *gin.Context)
	UpdatePhotoByID(ctx *gin.Context)
	DeletePhotoByID(ctx *gin.Context)
}

type photoController struct {
	photoService services.PhotoService
}

func NewPhotoController(ps services.PhotoService) PhotoController {
	return &photoController{
		photoService: ps,
	}
}

// @Create Photo
// @tag.name Photo
// @Description API for Adding Photo
// @Accept json
// @Produce json
// @Tags Photo
// @param Body body params.CreatePhotoRequest true "Create Photo"
// @Success 200 {object} params.CreatePhotoResponse
// @Router /photos [POST]
func (pc *photoController) CreatePhoto(ctx *gin.Context) {
	request := params.CreatePhotoRequest{}
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = helpers.ValidatePhotoRequest(request.Title, request.PhotoUrl)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := int(ctx.Keys["id"].(float64))
	res, err := pc.photoService.CreatePhoto(userID, request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

// @Get Photo
// @tag.name Photo
// @Description API for Get All Photo
// @Produce json
// @Tags Photo
// @Success 200 {object} []params.GetPhotoResponse
// @Router /photos [GET]
func (pc *photoController) GetPhotos(ctx *gin.Context) {
	res, err := pc.photoService.GetPhotosByUserID()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// @Update Photo
// @tag.name Photo
// @Description API for Updating Photo
// @Accept json
// @Produce json
// @Tags Photo
// @param Body body params.UpdatePhotoRequest true "Update Photo"
// @Success 200 {object} params.UpdatePhotoResponse
// @Router /photos/:photoId [PUT]
func (pc *photoController) UpdatePhotoByID(ctx *gin.Context) {
	IDStr := ctx.Param("photoid")
	ID, err := strconv.Atoi(IDStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var request params.UpdatePhotoRequest
	err = ctx.ShouldBind(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = helpers.ValidatePhotoRequest(request.Title, request.PhotoUrl)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := int(ctx.Keys["id"].(float64))
	res, err := pc.photoService.UpdatePhotoByID(ID, userID, request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)

}

// @Delete Photo
// @tag.name Photo
// @Description API for Deleting Photo
// @Produce json
// @Tags Photo
// @Success 200 {object} params.DeletePhotoResponse
// @Router /photos/:photoId [DELETE]
func (pc *photoController) DeletePhotoByID(ctx *gin.Context) {
	IDStr := ctx.Param("photoid")
	ID, err := strconv.Atoi(IDStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := int(ctx.Keys["id"].(float64))
	res, err := pc.photoService.DeletePhotoByID(ID, userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
