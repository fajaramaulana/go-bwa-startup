package user

import "time"

type UserFormatterRegister struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Occupation string `json:"occupation"`
	Token      string `json:"token"`
}

type UserFormaterReturnLogin struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Occupation     string    `json:"occupation"`
	AvatarFileName string    `json:"avatar_file_name"`
	Role           int       `json:"role"`
	Token          string    `json:"token"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func FormatUserRegister(user User, token string) UserFormatterRegister {
	formatter := UserFormatterRegister{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Occupation: user.Occupation,
		Token:      token,
	}

	return formatter
}

func FormatUserLogin(user User) UserFormaterReturnLogin {
	formatter := UserFormaterReturnLogin{
		ID:             user.ID,
		Name:           user.Name,
		Email:          user.Email,
		Occupation:     user.Occupation,
		AvatarFileName: user.Occupation,
		Role:           user.Role,
		Token:          user.Token,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}

	return formatter
}
