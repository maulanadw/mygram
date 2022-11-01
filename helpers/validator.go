package helpers

import (
	"errors"
	"mygram/params"
	"net/mail"
)

func ValidateUserRegisterRequest(request params.RegisterUserRequest) error {
	if len(request.Username) == 0 {
		return errors.New("username is required")
	}
	if len(request.Password) == 0 {
		return errors.New("password is required")
	}
	if len(request.Password) < 6 {
		return errors.New("password is required to have 6 minimum characters")
	}

	if len(request.Email) == 0 {
		return errors.New("email is required")
	}

	_, err := mail.ParseAddress(request.Email)
	if err != nil {
		return errors.New("email format is invalid")
	}

	if request.Age == 0 {
		return errors.New("age is required")
	}

	if request.Age < 8 {
		return errors.New("to register this app you need at least 8 years old")
	}

	return nil
}

func ValidateUserUpdateRequest(request params.UpdateUserRequest) error {
	if len(request.Username) == 0 {
		return errors.New("username is required")
	}

	if len(request.Email) == 0 {
		return errors.New("email is required")
	}

	_, err := mail.ParseAddress(request.Email)
	if err != nil {
		return errors.New("email format is invalid")
	}

	return nil
}

func ValidateUserLoginRequest(request params.LoginUserRequest) error {
	if len(request.Email) == 0 {
		return errors.New("email is required")
	}

	_, err := mail.ParseAddress(request.Email)
	if err != nil {
		return errors.New("email format is invalid")
	}

	if len(request.Password) == 0 {
		return errors.New("password is required")
	}

	return nil
}

func ValidatePhotoRequest(title, photoUrl string) error {
	if len(photoUrl) == 0 {
		return errors.New("photo_url is required")
	}

	if len(title) == 0 {
		return errors.New("title is required")
	}

	return nil
}

func ValidateSocialMediaRequest(name, socialMediaUrl string) error {
	if len(socialMediaUrl) == 0 {
		return errors.New("social_media_url is required")
	}

	if len(name) == 0 {
		return errors.New("title is required")
	}

	return nil
}

func ValidateCommentRequest(message string) error {
	if len(message) == 0 {
		return errors.New("message is required")
	}

	return nil
}
