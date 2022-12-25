package activity

import (
	"errors"
	"fmt"
	"time"

	"gotodo/utils"
)

// interface
type IActivityService interface {
	CreateActivity(input ActivityCreateInput) (Activity, error)
	GetAllActivity() []Activity
	GetActivityByID(id int) (Activity, error)
	DeleteByID(id int) error
	UpdateByID(input ActivityUpdateInput, id int) (Activity, error)
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
	if err := utils.ValidateID(id); err != nil {
		return Activity{}, err
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

func (s *activityService) DeleteByID(id int) error {
	if err := utils.ValidateID(id); err != nil {
		return err
	}

	// find activity by id
	activity, err := s.activityRepo.FindByID(id)
	if err != nil {
		return err
	}

	if activity.ID == 0 {
		errMsg := fmt.Sprintf("Activity with ID %v Not Found", id)
		return errors.New(errMsg)
	}

	// call repo and return
	return s.activityRepo.DeleteByID(id)
}

func (s *activityService) UpdateByID(input ActivityUpdateInput, id int) (Activity, error) {
	if err := utils.ValidateID(id); err != nil {
		return Activity{}, err
	}

	// find activity by id
	activity, err := s.activityRepo.FindByID(id)
	if err != nil {
		return activity, err
	}

	if activity.ID == 0 {
		errMsg := fmt.Sprintf("Activity with ID %v Not Found", id)
		return activity, errors.New(errMsg)
	}

	// update
	activity.Title = input.Title
	activity.UpdatedAt = time.Now()

	// call repo
	activityUpdated, err := s.activityRepo.UpdateById(activity)
	if err != nil {
		return activityUpdated, err
	}

	return activityUpdated, nil
}
