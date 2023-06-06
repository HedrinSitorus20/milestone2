package accounts

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func (r Repository) Login(a *Actor, lr *LoginRequest) error {
	return r.db.Where("username = ?", lr.Username).First(&a).Error
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}
