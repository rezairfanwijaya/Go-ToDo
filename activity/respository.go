package activity

import "gorm.io/gorm"

// interface
type IActivityRepository interface {
	Save(activity Activity) (Activity, error)
	FindAll() []Activity
	FindByID(id int) (Activity, error)
	DeleteByID(id int) error
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

func (r *repository) FindAll() []Activity {
	var activities []Activity

	r.db.Find(&activities)

	return activities
}

func (r *repository) FindByID(id int) (Activity, error) {
	var activity Activity

	if err := r.db.Where("id = ?", id).Find(&activity).Error; err != nil {
		return activity, err
	}

	return activity, nil
}

func (r *repository) DeleteByID(id int) error {
	var activity Activity
	if err := r.db.Where("id = ?", id).Delete(&activity).Error; err != nil {
		return err
	}
	return nil
}
