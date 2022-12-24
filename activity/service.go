package activity

import "time"

// interface
type IActivityService interface {
	CreateActivity(input ActivityCreateInput) (Activity, error)
	GetAllActivity() []Activity
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
