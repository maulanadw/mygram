package services

import (
	"errors"
	"mygram/helpers"
	"mygram/models"
	"mygram/params"
	"mygram/repositories"
)

type UserService interface {
	Register(request params.RegisterUserRequest) (*params.RegisterUserResponse, error)
	Login(request params.LoginUserRequest) (*params.LoginUserResponse, error)
	UpdateUserByID(ID int, request params.UpdateUserRequest) (*params.UpdateUserResponse, error)
	DeleteUserByID(ID int) (*params.DeleteUserResponse, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(ur repositories.UserRepository) UserService {
	return &userService{
		userRepo: ur,
	}
}

func (us *userService) Register(request params.RegisterUserRequest) (*params.RegisterUserResponse, error) {
	hashedPassword, err := helpers.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	userModel := models.User{
		Username: request.Username,
		Password: hashedPassword,
		Email:    request.Email,
		Age:      request.Age,
	}

	user, err := us.userRepo.CreateUser(userModel)
	if err != nil {
		return nil, err
	}

	registerResponse := params.RegisterUserResponse{
		ID:       user.ID,
		Age:      user.Age,
		Email:    user.Email,
		Username: user.Username,
	}
	return &registerResponse, nil

}

func (us *userService) Login(request params.LoginUserRequest) (*params.LoginUserResponse, error) {
	user, err := us.userRepo.GetUserByEmail(request.Email)
	if err != nil {
		return nil, err
	}

	isValid := helpers.ValidatePassword(request.Password, user.Password)
	if !isValid {
		err = errors.New("password invalid")
		return nil, err
	}

	token, err := helpers.GenerateToken(int(user.ID), user.Email)
	if err != nil {
		return nil, err
	}

	loginResponse := params.LoginUserResponse{
		Token: token,
	}
	return &loginResponse, nil
}

func (us *userService) UpdateUserByID(ID int, request params.UpdateUserRequest) (*params.UpdateUserResponse, error) {
	userModel := models.User{
		Username: request.Username,
		Email:    request.Email,
	}

	res, err := us.userRepo.UpdateUserByID(ID, userModel)
	if err != nil {
		return nil, err
	}

	updateUserResponse := params.UpdateUserResponse{
		ID:        res.ID,
		Age:       res.Age,
		Email:     res.Email,
		Username:  res.Username,
		UpdatedAt: res.UpdatedAt,
	}
	return &updateUserResponse, nil
}

func (us *userService) DeleteUserByID(ID int) (*params.DeleteUserResponse, error) {
	err := us.userRepo.DeleteUserByID(ID)
	if err != nil {
		return nil, err
	}

	deleteUserResponse := params.DeleteUserResponse{
		Message: "Your account has been successfully deleted",
	}

	return &deleteUserResponse, nil
}
