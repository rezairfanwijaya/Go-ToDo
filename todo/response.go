package todo

import "time"

type todoCreateFormatter struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	Title      string    `json:"title"`
	ActivityID int       `json:"activity_group_id"`
	IsActive   bool      `json:"is_active"`
	Priority   string    `json:"priority"`
	CreatredAt time.Time `json:"updatedAt"`
	UpdatedAt  time.Time `json:"createdAt"`
}

type todoFormatter struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	ActivityID int       `json:"activity_group_id"`
	Title      string    `json:"title"`
	IsActive   bool      `json:"is_active"`
	Priority   string    `json:"priority"`
	CreatredAt time.Time `json:"updatedAt"`
	UpdatedAt  time.Time `json:"createdAt"`
}

type TodoAfterDelete struct{}

func FormatterCreateTodo(todo Todo) todoCreateFormatter {
	return todoCreateFormatter{
		ID:         todo.ID,
		ActivityID: todo.ActivityID,
		Title:      todo.Title,
		IsActive:   todo.IsActive,
		Priority:   todo.Priority,
		CreatredAt: todo.CreatredAt,
		UpdatedAt:  todo.UpdatedAt,
	}
}

func FormatterGetTodo(todo Todo) todoFormatter {
	return todoFormatter{
		ID:         todo.ID,
		ActivityID: todo.ActivityID,
		Title:      todo.Title,
		IsActive:   todo.IsActive,
		Priority:   todo.Priority,
		CreatredAt: todo.CreatredAt,
		UpdatedAt:  todo.UpdatedAt,
	}
}

func FormatterTodos(todos []Todo) []todoFormatter {
	var res []todoFormatter

	for _, todo := range todos {
		res = append(res, FormatterGetTodo(todo))
	}

	return res
}
