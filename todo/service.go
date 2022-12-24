package todo

import (
	"gotodo/activity"
	"gotodo/utils"
	"time"
)

// intrerface
type ITodoService interface {
	CreateTodo(input TodoCreateInput) (Todo, error)
}

type todoService struct {
	repoTodo        ITodoRepository
	serviceActivity activity.IActivityService
}

// new service
func NewTodoService(repoTodo ITodoRepository, serviceActivity activity.IActivityService) *todoService {
	return &todoService{repoTodo, serviceActivity}
}

func (s *todoService) CreateTodo(input TodoCreateInput) (Todo, error) {
	if err := utils.ValidateID(input.ActivityID); err != nil {
		return Todo{}, err
	}

	// activity id must be exist
	activity, err := s.serviceActivity.GetActivityByID(input.ActivityID)
	if err != nil {
		return Todo{}, err
	}

	// binding
	var todo Todo
	todo.Title = input.Title
	todo.ActivityID = activity.ID
	todo.IsActive = input.IsActive
	todo.Priority = "very-high"
	todo.CreatredAt = time.Now()
	todo.UpdatedAt = time.Now()

	// call repo
	todoSaved, err := s.repoTodo.Save(todo)
	if err != nil {
		return todoSaved, err
	}

	return todoSaved, nil
}
