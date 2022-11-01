package services

import (
	"mygram/models"
	"mygram/params"
	"mygram/repositories"
)

type SocialMediaService interface {
	CreateSocialMedia(userID int, request params.CreateSocialMediaRequest) (*params.CreateSocialMediaResponse, error)
	GetSocialMedias() (*params.GetSocialMediaResponses, error)
	UpdateSocialMediaByID(ID int, userID int, request params.UpdateSocialMediaRequest) (*params.UpdateSocialMediaResponse, error)
	DeleteSocialMediaByID(ID, userID int) (*params.DeleteSocialMediaResponse, error)
}

type socialMediaService struct {
	socialMediaRepo repositories.SocialMediaRepository
}

func NewSocialMediaService(smr repositories.SocialMediaRepository) SocialMediaService {
	return &socialMediaService{
		socialMediaRepo: smr,
	}
}

func toSocialMediaParam(socialMediaModel models.SocialMedia) params.SocialMedia {
	user := params.User{
		ID:       &socialMediaModel.UserID,
		Username: socialMediaModel.User.Username,
		Email:    socialMediaModel.User.Email,
	}
	return params.SocialMedia{
		ID:             socialMediaModel.ID,
		Name:           socialMediaModel.Name,
		UserID:         socialMediaModel.UserID,
		CreatedAt:      socialMediaModel.CreatedAt,
		UpdatedAt:      socialMediaModel.UpdatedAt,
		SocialMediaUrl: socialMediaModel.SocialMediaUrl,
		User:           user,
	}
}

func toSocialMediaParams(socialMediaModels []models.SocialMedia) []params.SocialMedia {
	socialMediaParams := make([]params.SocialMedia, len(socialMediaModels))
	for idx, socialMediaModel := range socialMediaModels {
		socialMediaParams[idx] = toSocialMediaParam(socialMediaModel)
	}

	return socialMediaParams
}

func (sms *socialMediaService) CreateSocialMedia(userID int, request params.CreateSocialMediaRequest) (*params.CreateSocialMediaResponse, error) {
	socialMedia := models.SocialMedia{
		Name:           request.Name,
		SocialMediaUrl: request.SocialMediaUrl,
		UserID:         userID,
	}

	res, err := sms.socialMediaRepo.CreateSocialMedia(socialMedia)
	if err != nil {
		return nil, err
	}

	response := params.CreateSocialMediaResponse{
		ID:             res.ID,
		Name:           res.Name,
		SocialMediaUrl: res.SocialMediaUrl,
		UserID:         res.UserID,
		CreatedAt:      res.CreatedAt,
	}

	return &response, nil

}

func (sms *socialMediaService) GetSocialMedias() (*params.GetSocialMediaResponses, error) {
	res, err := sms.socialMediaRepo.GetSocialMedias()
	if err != nil {
		return nil, err
	}

	getSocialMediaResponse := params.GetSocialMediaResponses{
		SocialMedias: toSocialMediaParams(res),
	}

	return &getSocialMediaResponse, nil

}

func (sms *socialMediaService) UpdateSocialMediaByID(ID int, userID int, request params.UpdateSocialMediaRequest) (*params.UpdateSocialMediaResponse, error) {
	socialMedia := models.SocialMedia{
		Name:           request.Name,
		SocialMediaUrl: request.SocialMediaUrl,
	}

	res, err := sms.socialMediaRepo.UpdateSocialMediaByID(ID, userID, socialMedia)
	if err != nil {
		return nil, err
	}

	updateSocialMediaResponse := params.UpdateSocialMediaResponse{
		ID:             res.ID,
		Name:           res.Name,
		SocialMediaUrl: res.SocialMediaUrl,
		UserID:         res.UserID,
		UpdatedAt:      res.UpdatedAt,
	}

	return &updateSocialMediaResponse, nil
}

func (sms *socialMediaService) DeleteSocialMediaByID(ID, userID int) (*params.DeleteSocialMediaResponse, error) {
	err := sms.socialMediaRepo.DeleteSocialMediaByID(ID, userID)
	if err != nil {
		return nil, err
	}

	deleteSocialMediaResponse := params.DeleteSocialMediaResponse{
		Message: "Your social media has been successfully deleted",
	}

	return &deleteSocialMediaResponse, nil
}
