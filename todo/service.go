package todo

import (
	"errors"
	"fmt"
	"gotodo/activity"
	"gotodo/utils"
	"time"
)

// intrerface
type ITodoService interface {
	CreateTodo(input TodoCreateInput) (Todo, error)
	GetTodoById(id int) (Todo, error)
	GetAllTodo(id int, isHaveQuery bool) ([]Todo, error)
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

func (s *todoService) GetTodoById(id int) (Todo, error) {
	if err := utils.ValidateID(id); err != nil {
		return Todo{}, err
	}

	// call repo
	todo, err := s.repoTodo.FindByID(id)
	if err != nil {
		return todo, err
	}

	if todo.ID == 0 {
		errMsg := fmt.Sprintf("Todo with ID %v Not Found", id)
		return todo, errors.New(errMsg)
	}

	return todo, nil
}

func (s *todoService) GetAllTodo(id int, isHaveQuery bool) ([]Todo, error) {
	if isHaveQuery {
		if err := utils.ValidateID(id); err != nil {
			return []Todo{}, err
		}

		// call repo
		todos, err := s.repoTodo.FindByActivityID(id)
		if err != nil {
			return todos, err
		}

		if len(todos) == 0 {
			errMsg := fmt.Sprintf("Activity Group with ID %v Not Found", id)
			return todos, errors.New(errMsg)
		}

		return todos, nil
	}

	// call repo
	return s.repoTodo.FindAll(), nil
}
