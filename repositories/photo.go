package repositories

import (
	"errors"
	"mygram/models"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	CreatePhoto(photo models.Photo) (*models.Photo, error)
	GetPhotosByUserID() ([]models.Photo, error)
	UpdatePhotoByID(ID int, photo models.Photo) (*models.Photo, error)
	DeletePhotoByID(ID int, userID int) error
}

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	db.AutoMigrate(&models.Photo{})
	return &photoRepository{
		db: db,
	}
}

func (pr *photoRepository) CreatePhoto(photo models.Photo) (*models.Photo, error) {
	err := pr.db.Create(&photo).Error
	return &photo, err
}

func (pr *photoRepository) GetPhotosByUserID() ([]models.Photo, error) {
	var photos []models.Photo
	err := pr.db.Preload("User").Find(&photos).Error

	return photos, err

}

func (pr *photoRepository) UpdatePhotoByID(ID int, photo models.Photo) (*models.Photo, error) {
	var result models.Photo
	err := pr.db.Transaction(func(tx *gorm.DB) error {
		err := tx.First(&result, ID).Error
		if err != nil {
			return err
		}

		if uint(result.UserID) != uint(photo.UserID) {
			err := errors.New("you're forbidden to update this data")
			return err
		}

		err = tx.Where("id=?", ID).Updates(&photo).Find(&result).Error
		if err != nil {
			return err
		}

		return nil
	})

	return &result, err
}

func (pr *photoRepository) DeletePhotoByID(ID int, userID int) error {
	var result models.Photo
	err := pr.db.Transaction(func(tx *gorm.DB) error {
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
