package todo

import "time"

type todoFormatter struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	ActivityID int       `json:"activity_group_id"`
	Title      string    `json:"title"`
	IsActive   bool      `json:"is_active"`
	Priority   string    `json:"priority"`
	CreatredAt time.Time `json:"updatedAt"`
	UpdatedAt  time.Time `json:"createdAt"`
}

func FormatterTodo(todo Todo) todoFormatter {
	return todoFormatter{
		ID:         todo.ID,
		ActivityID: todo.ActivityID,
		Title:      todo.Title,
		IsActive:   todo.IsActive,
		Priority:   todo.Priority,
		CreatredAt: todo.CreatredAt,
		UpdatedAt:  todo.CreatredAt,
	}
}

func FormatterTodos(todos []Todo) []todoFormatter {
	var res []todoFormatter

	for _, todo := range todos {
		res = append(res, FormatterTodo(todo))
	}

	return res
}
