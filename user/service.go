package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	LoginUser(input LoginUserInput) (User, error)
	IsEmailAvailable(input CheckEmailAvailability) (bool, error)
	UploadAvatar(ID int, fileLocation string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	// mapping struct input ke struct user
	// simpan struct user melalui repository
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)
	user.Role = 2

	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) LoginUser(input LoginUserInput) (User, error) {
	userInput := User{}
	userInput.Email = input.Email
	userInput.PasswordHash = input.Password

	userLogged, err := s.repository.FindByEmail(userInput)

	if userLogged.ID == 0 {
		return userLogged, errors.New("User Not Found")
	}

	if err != nil {
		return userLogged, errors.New("User Not Found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userLogged.PasswordHash), []byte(userInput.PasswordHash))

	if err != nil {
		return userLogged, errors.New("Incorrect Password")
	}

	return userLogged, nil
}

func (s *service) IsEmailAvailable(input CheckEmailAvailability) (bool, error) {
	userInput := User{}
	userInput.Email = input.Email

	emailExist, err := s.repository.FindByEmail(userInput)

	if err != nil {
		return false, err
	}

	if emailExist.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) UploadAvatar(ID int, fileLocation string) (User, error) {
	// dapatkan user berdasarkan id

	user, err := s.repository.FindById(ID)

	if err != nil {
		return user, err
	}
	// update attribute avatar file name
	user.AvatarFileName = fileLocation

	// simpan perubahan avatar file name
	updatedUser, err := s.repository.Update(user)

	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}
