package repositories

import (
	"mygram/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUserByID(ID int, user models.User) (*models.User, error)
	DeleteUserByID(ID int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	db.AutoMigrate(&models.User{})
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) CreateUser(user models.User) (models.User, error) {
	err := ur.db.Create(&user).Error
	return user, err
}

func (ur *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := ur.db.Where("email=?", email).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) UpdateUserByID(ID int, user models.User) (*models.User, error) {
	var result models.User
	err := ur.db.Transaction(func(tx *gorm.DB) error {
		err := tx.First(&result, ID).Error
		if err != nil {
			return err
		}

		err = tx.Where("id=?", ID).Updates(&user).Find(&result).Error
		if err != nil {
			return err
		}

		return nil
	})

	return &result, err
}

func (ur *userRepository) DeleteUserByID(ID int) error {
	var user models.User
	err := ur.db.Transaction(func(tx *gorm.DB) error {
		err := tx.First(&user, ID).Error
		if err != nil {
			return err
		}

		err = tx.Where("id=?", ID).Delete(&user).Error
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
