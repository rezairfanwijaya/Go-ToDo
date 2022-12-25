package activity

import "gorm.io/gorm"

// interface
type IActivityRepository interface {
	Save(activity ActivityGroup) (ActivityGroup, error)
	FindAll() []ActivityGroup
	FindByID(id int) (ActivityGroup, error)
	DeleteByID(id int) error
	UpdateById(activity ActivityGroup) (ActivityGroup, error)
}

type repository struct {
	db *gorm.DB
}

// new repo
func NewActivityRespository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(activity ActivityGroup) (ActivityGroup, error) {
	if err := r.db.Create(&activity).Error; err != nil {
		return activity, err
	}

	return activity, nil
}

func (r *repository) FindAll() []ActivityGroup {
	var activities []ActivityGroup

	r.db.Find(&activities)

	return activities
}

func (r *repository) FindByID(id int) (ActivityGroup, error) {
	var activity ActivityGroup

	if err := r.db.Where("id = ?", id).Find(&activity).Error; err != nil {
		return activity, err
	}

	return activity, nil
}

func (r *repository) DeleteByID(id int) error {
	var activity ActivityGroup

	if err := r.db.Where("id = ?", id).Delete(&activity).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateById(activity ActivityGroup) (ActivityGroup, error) {
	if err := r.db.Save(&activity).Error; err != nil {
		return activity, err
	}

	return activity, nil
}
