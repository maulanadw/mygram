package repositories

import (
	"errors"
	"fmt"
	"mygram/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	CreateComment(request models.Comment) (*models.Comment, error)
	GetCommentsByUserID() ([]models.Comment, error)
	UpdateCommentByID(ID int, userID int, request models.Comment) (*models.Comment, error)
	DeleteCommentByID(ID, userID int) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	db.AutoMigrate(models.Comment{})
	return &commentRepository{
		db: db,
	}
}

func (cr *commentRepository) CreateComment(comment models.Comment) (*models.Comment, error) {
	err := cr.db.Create(&comment).Error
	return &comment, err
}

func (cr *commentRepository) GetCommentsByUserID() ([]models.Comment, error) {
	var comments []models.Comment
	err := cr.db.Preload("User").Preload("Photo").Find(&comments).Error

	return comments, err
}

func (cr *commentRepository) UpdateCommentByID(ID int, userID int, request models.Comment) (*models.Comment, error) {
	var result models.Comment
	err := cr.db.Transaction(func(tx *gorm.DB) error {
		err := tx.First(&result, ID).Error
		if err != nil {
			return err
		}

		if result.UserID != userID {
			err := errors.New("you're forbidden to update this data")
			return err
		}

		err = tx.Where("id=?", ID).Updates(&request).Error
		if err != nil {
			fmt.Println("2")
			return err
		}

		err = cr.db.Preload("Photo").First(&result, ID).Error
		if err != nil {
			return err
		}

		return nil
	})

	return &result, err
}

func (cr *commentRepository) DeleteCommentByID(ID, userID int) error {
	var result models.Comment
	err := cr.db.Transaction(func(tx *gorm.DB) error {
		err := tx.First(&result).Error
		if err != nil {
			return err
		}

		if result.UserID != userID {
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
