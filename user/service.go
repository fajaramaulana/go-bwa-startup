package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	LoginUser(input LoginUserInput) (User, error)
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
		return userLogged, errors.New("User Not Found1")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userLogged.PasswordHash), []byte(userInput.PasswordHash))

	if err != nil {
		return userLogged, errors.New("Incorrect Password")
	}

	return userLogged, nil
}
