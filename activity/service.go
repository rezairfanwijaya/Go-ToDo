package activity

import (
	"errors"
	"fmt"
	"time"
)

// interface
type IActivityService interface {
	CreateActivity(input ActivityCreateInput) (Activity, error)
	GetAllActivity() []Activity
	GetActivityByID(id int) (Activity, error)
}

type activityService struct {
	activityRepo IActivityRepository
}

// new service
func NewActivityService(activityRepo IActivityRepository) *activityService {
	return &activityService{activityRepo}
}

func (s *activityService) CreateActivity(input ActivityCreateInput) (Activity, error) {
	// bind
	var activity Activity
	activity.Email = input.Email
	activity.Title = input.Title
	activity.CreatredAt = time.Now()
	activity.UpdatedAt = time.Now()

	// call repo
	newActivity, err := s.activityRepo.Save(activity)
	if err != nil {
		return newActivity, err
	}

	return newActivity, nil
}

func (s *activityService) GetAllActivity() []Activity {
	return s.activityRepo.FindAll()
}

func (s *activityService) GetActivityByID(id int) (Activity, error) {
	// validation
	if id <= 0 {
		errMsg := fmt.Sprintf("id must grather then 0")
		return Activity{}, errors.New(errMsg)
	}

	// call service
	activity, err := s.activityRepo.FindByID(id)
	if err != nil {
		return activity, err
	}

	// activity id not found
	if activity.ID == 0 {
		errMsg := fmt.Sprintf("Activity with ID %v Not Found", id)
		return activity, errors.New(errMsg)
	}

	return activity, nil
}
