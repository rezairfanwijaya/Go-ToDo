package activity

import "gorm.io/gorm"

// interface
type IActivityRepository interface {
	Save(activity Activity) (Activity, error)
}

type repository struct {
	db *gorm.DB
}

// new repo
func NewActivityRespository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(activity Activity) (Activity, error) {
	if err := r.db.Create(&activity).Error; err != nil {
		return activity, err
	}

	return activity, nil
}
