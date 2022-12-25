package activity

import (
	"errors"
	"fmt"
	"time"

	"gotodo/utils"
)

// interface
type IActivityService interface {
	CreateActivity(input ActivityCreateInput) (ActivityGroup, error)
	GetAllActivity() []ActivityGroup
	GetActivityByID(id int) (ActivityGroup, error)
	DeleteByID(id int) error
	UpdateByID(input ActivityUpdateInput, id int) (ActivityGroup, error)
}

type activityService struct {
	activityRepo IActivityRepository
}

// new service
func NewActivityService(activityRepo IActivityRepository) *activityService {
	return &activityService{activityRepo}
}

func (s *activityService) CreateActivity(input ActivityCreateInput) (ActivityGroup, error) {
	// bind
	var activity ActivityGroup
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

func (s *activityService) GetAllActivity() []ActivityGroup {
	return s.activityRepo.FindAll()
}

func (s *activityService) GetActivityByID(id int) (ActivityGroup, error) {
	// validation
	if err := utils.ValidateID(id); err != nil {
		return ActivityGroup{}, err
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

func (s *activityService) UpdateByID(input ActivityUpdateInput, id int) (ActivityGroup, error) {
	if err := utils.ValidateID(id); err != nil {
		return ActivityGroup{}, err
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
