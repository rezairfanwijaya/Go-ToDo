package activity

import "gorm.io/gorm"

// interface
type IRepository interface {
}

type repository struct {
	db *gorm.DB
}

// new repo
func NewRespository(db *gorm.DB) *repository {
	return &repository{db}
}
