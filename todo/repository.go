package todo

import "gorm.io/gorm"

// interface
type ITodoRepository interface {
	Save(todo Todo) (Todo, error)
	FindByID(id int) (Todo, error)
}

type todoRepository struct {
	db *gorm.DB
}

// new repo
func NewTodoRepository(db *gorm.DB) *todoRepository {
	return &todoRepository{db}
}

func (r *todoRepository) Save(todo Todo) (Todo, error) {
	if err := r.db.Create(&todo).Error; err != nil {
		return todo, err
	}

	return todo, nil
}

func (r *todoRepository) FindByID(id int) (Todo, error) {
	var todo Todo
	if err := r.db.Where("id = ?", id).Find(&todo).Error; err != nil {
		return todo, err
	}

	return todo, nil
}
