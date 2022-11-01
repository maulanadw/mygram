package repositories

import (
	"errors"
	"mygram/models"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	CreateSocialMedia(socialMedia models.SocialMedia) (*models.SocialMedia, error)
	GetSocialMedias() ([]models.SocialMedia, error)
	UpdateSocialMediaByID(ID int, userID int, socialMedia models.SocialMedia) (*models.SocialMedia, error)
	DeleteSocialMediaByID(ID int, userID int) error
}

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepository {
	db.AutoMigrate(&models.SocialMedia{})
	return &socialMediaRepository{
		db: db,
	}
}

func (smr *socialMediaRepository) CreateSocialMedia(socialMedia models.SocialMedia) (*models.SocialMedia, error) {
	err := smr.db.Create(&socialMedia).Error
	return &socialMedia, err
}

func (smr *socialMediaRepository) GetSocialMedias() ([]models.SocialMedia, error) {
	var socialMedias []models.SocialMedia
	err := smr.db.Preload("User").Find(&socialMedias).Error
	return socialMedias, err
}

func (smr *socialMediaRepository) UpdateSocialMediaByID(ID int, userID int, socialMedia models.SocialMedia) (*models.SocialMedia, error) {
	var result models.SocialMedia
	err := smr.db.Transaction(func(tx *gorm.DB) error {
		err := tx.First(&result, ID).Error
		if err != nil {
			return err
		}

		if uint(result.UserID) != uint(userID) {
			err := errors.New("you're forbidden to delete this data")
			return err
		}

		err = tx.Where("id=?", ID).Updates(&socialMedia).Find(&result).Error
		if err != nil {
			return err
		}
		return nil
	})

	return &result, err
}

func (smr *socialMediaRepository) DeleteSocialMediaByID(ID int, userID int) error {
	var result models.SocialMedia
	err := smr.db.Transaction(func(tx *gorm.DB) error {
		err := tx.First(&result, ID).Error
		if err != nil {
			return err
		}

		if uint(result.UserID) != uint(userID) {
			err := errors.New("you're forbidden to delete this data")
			return err
		}

		err = tx.Where("id=?", ID).Delete(&result).Error
		if err != nil {
			return err
		}
		return nil
	})

	return err
}
