package user

import "gorm.io/gorm"

type Repository interface {
	Save(u User) (User, error)
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
