package user

import "gorm.io/gorm"

type Repository interface {
	Save(u User) (User, error)
	FindByEmail(uInput User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(u User) (User, error) {
	err := r.db.Create(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

func (r *repository) FindByEmail(uInput User) (User, error) {
	var user User
	err := r.db.Where("email = ?", uInput.Email).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
