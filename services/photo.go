package services

import (
	"mygram/models"
	"mygram/params"
	"mygram/repositories"
)

type PhotoService interface {
	CreatePhoto(userID int, request params.CreatePhotoRequest) (*params.CreatePhotoResponse, error)
	GetPhotosByUserID() ([]params.GetPhotoResponse, error)
	UpdatePhotoByID(ID int, userID int, request params.UpdatePhotoRequest) (*params.UpdatePhotoResponse, error)
	DeletePhotoByID(ID int, userID int) (*params.DeletePhotoResponse, error)
}

type photoService struct {
	photoRepo repositories.PhotoRepository
}

func NewPhotoService(pr repositories.PhotoRepository) PhotoService {
	return &photoService{
		photoRepo: pr,
	}
}

func toGetPhotoResponse(photoModel models.Photo) params.GetPhotoResponse {
	userParam := params.User{
		Username: photoModel.User.Username,
		Email:    photoModel.User.Email,
	}

	return params.GetPhotoResponse{
		ID:        int(photoModel.ID),
		Title:     photoModel.Title,
		Caption:   photoModel.Caption,
		PhotoUrl:  photoModel.PhotoUrl,
		User:      userParam,
		CreatedAt: photoModel.CreatedAt,
		UpdatedAt: photoModel.UpdatedAt,
	}
}

func toGetPhotoResponses(photoModels []models.Photo) []params.GetPhotoResponse {
	getPhotoResponses := make([]params.GetPhotoResponse, len(photoModels))
	for idx, photoModel := range photoModels {
		getPhotoResponses[idx] = toGetPhotoResponse(photoModel)
	}

	return getPhotoResponses
}

func (ps *photoService) CreatePhoto(userID int, request params.CreatePhotoRequest) (*params.CreatePhotoResponse, error) {
	photo := models.Photo{
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoUrl: request.PhotoUrl,
		UserID:   userID,
	}

	res, err := ps.photoRepo.CreatePhoto(photo)
	if err != nil {
		return nil, err
	}

	createPhotoResponse := params.CreatePhotoResponse{
		ID:        int(res.ID),
		Title:     res.Title,
		Caption:   res.Caption,
		PhotoUrl:  res.PhotoUrl,
		UserID:    res.UserID,
		CreatedAt: res.CreatedAt,
	}
	return &createPhotoResponse, err
}
func (ps *photoService) GetPhotosByUserID() ([]params.GetPhotoResponse, error) {
	result, err := ps.photoRepo.GetPhotosByUserID()
	if err != nil {
		return nil, err
	}

	return toGetPhotoResponses(result), nil
}
func (ps *photoService) UpdatePhotoByID(ID int, userID int, request params.UpdatePhotoRequest) (*params.UpdatePhotoResponse, error) {
	photoModel := models.Photo{
		Title:    request.Title,
		PhotoUrl: request.PhotoUrl,
		Caption:  request.Caption,
		UserID:   userID,
	}

	res, err := ps.photoRepo.UpdatePhotoByID(ID, photoModel)
	if err != nil {
		return nil, err
	}

	photoParam := params.UpdatePhotoResponse{
		ID:        int(res.ID),
		Title:     res.Title,
		Caption:   res.Caption,
		PhotoUrl:  res.PhotoUrl,
		UserID:    res.UserID,
		UpdatedAt: res.UpdatedAt,
	}

	return &photoParam, nil

}
func (ps *photoService) DeletePhotoByID(ID int, userID int) (*params.DeletePhotoResponse, error) {
	err := ps.photoRepo.DeletePhotoByID(ID, userID)
	if err != nil {
		return nil, err
	}

	deletePhotoResponse := params.DeletePhotoResponse{
		Message: "Your photo has been successfully deleted",
	}

	return &deletePhotoResponse, nil
}
