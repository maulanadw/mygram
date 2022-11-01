package controllers

import (
	"mygram/helpers"
	"mygram/params"
	"mygram/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	UpdateUserByID(ctx *gin.Context)
	DeleteUserByID(ctx *gin.Context)
}

type userContorller struct {
	userService services.UserService
}

func NewUserController(us services.UserService) UserController {
	return &userContorller{
		userService: us,
	}
}

// @Register User
// @tag.name User
// @Description API for user registration
// @Accept json
// @Produce json
// @Tags User
// @param Body body params.RegisterUserRequest true "Register User"
// @Success 200 {object} params.RegisterUserResponse
// @Router /users/register [POST]
func (uc *userContorller) Register(ctx *gin.Context) {
	user := params.RegisterUserRequest{}
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = helpers.ValidateUserRegisterRequest(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := uc.userService.Register(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

// @Login User
// @tag.name User
// @Description API for user login
// @Accept json
// @Produce json
// @Tags User
// @param Body body params.LoginUserRequest true "Login user"
// @Success 200 {object} params.LoginUserResponse
// @Router /users/login [POST]
func (uc *userContorller) Login(ctx *gin.Context) {
	user := params.LoginUserRequest{}
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = helpers.ValidateUserLoginRequest(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := uc.userService.Login(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// @Update User
// @tag.name User
// @Description API for update user
// @Accept json
// @Produce json
// @Tags User
// @param Body body params.UpdateUserRequest true "Update user"
// @Success 200 {object} params.UpdateUserResponse
// @Router /users/:userid [PUT]
func (uc *userContorller) UpdateUserByID(ctx *gin.Context) {
	user := params.UpdateUserRequest{}
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = helpers.ValidateUserUpdateRequest(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	IDstr := ctx.Param("id")
	ID, err := strconv.Atoi(IDstr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := int(ctx.Keys["id"].(float64))
	if ID != userID {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "you're forbidden to update this data",
		})
		return
	}

	res, err := uc.userService.UpdateUserByID(ID, user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)

}

// @Delete User
// @tag.name User
// @Description API for delete user
// @Accept json
// @Produce json
// @Tags User
// @Success 200 {object} params.DeleteUserResponse
// @Router /users/:userid [DELETE]
func (uc *userContorller) DeleteUserByID(ctx *gin.Context) {
	IDstr := ctx.Param("id")
	ID, err := strconv.Atoi(IDstr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := int(ctx.Keys["id"].(float64))
	if ID != userID {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "you're forbidden to delete this data",
		})
		return
	}

	res, err := uc.userService.DeleteUserByID(ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
