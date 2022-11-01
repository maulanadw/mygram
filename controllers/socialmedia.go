package controllers

import (
	"mygram/helpers"
	"mygram/params"
	"mygram/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SocialMediaController interface {
	CreateSocialMedia(ctx *gin.Context)
	GetSocialMedias(ctx *gin.Context)
	UpdateSocialMediaByID(ctx *gin.Context)
	DeleteSocialMediaByID(ctx *gin.Context)
}

type socialMediaController struct {
	socialMediaService services.SocialMediaService
}

func NewSocialMediaController(sms services.SocialMediaService) SocialMediaController {
	return &socialMediaController{
		socialMediaService: sms,
	}
}

// @Create Social Media
// @tag.name Social MEdia
// @Description API for Adding Social Media
// @Accept json
// @Produce json
// @Tags Social Media
// @param Body body params.CreateSocialMediaRequest true "Create Social Media"
// @Success 200 {object} params.CreateSocialMediaResponse
// @Router /socialmedias [POST]
func (smc *socialMediaController) CreateSocialMedia(ctx *gin.Context) {
	var request params.CreateSocialMediaRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = helpers.ValidateSocialMediaRequest(request.Name, request.SocialMediaUrl)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := int(ctx.Keys["id"].(float64))
	res, err := smc.socialMediaService.CreateSocialMedia(userID, request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

// @Get Social Media
// @tag.name Social Media
// @Description API for Get All Social Media
// @Produce json
// @Tags Social Media
// @Success 200 {object} params.GetSocialMediaResponses
// @Router /socialmedias [GET]
func (smc *socialMediaController) GetSocialMedias(ctx *gin.Context) {
	res, err := smc.socialMediaService.GetSocialMedias()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// @Update Social Media
// @tag.name Social MEdia
// @Description API for Updating Social Media
// @Accept json
// @Produce json
// @Tags Social Media
// @param Body body params.UpdateSocialMediaRequest true "Update Social Media"
// @Success 200 {object} params.UpdateSocialMediaResponse
// @Router /socialmedias/:socialMediaId [PUT]
func (smc *socialMediaController) UpdateSocialMediaByID(ctx *gin.Context) {
	IDStr := ctx.Param("socialMediaId")
	ID, err := strconv.Atoi(IDStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var request params.UpdateSocialMediaRequest
	err = ctx.ShouldBind(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = helpers.ValidateSocialMediaRequest(request.Name, request.SocialMediaUrl)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := int(ctx.Keys["id"].(float64))
	res, err := smc.socialMediaService.UpdateSocialMediaByID(ID, userID, request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)

}

// @Delete Social Media
// @tag.name Social Media
// @Description API for Deleting Social Media
// @Produce json
// @Tags Social Media
// @Success 200 {object} params.DeleteSocialMediaResponse
// @Router /socialmedias/:socialMediaId [DELETE]
func (smc *socialMediaController) DeleteSocialMediaByID(ctx *gin.Context) {
	IDStr := ctx.Param("socialMediaId")
	ID, err := strconv.Atoi(IDStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := int(ctx.Keys["id"].(float64))
	res, err := smc.socialMediaService.DeleteSocialMediaByID(ID, userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
