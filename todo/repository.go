package todo

import "gorm.io/gorm"

// interface
type ITodoRepository interface {
	Save(todo Todo) (Todo, error)
	FindByID(id int) (Todo, error)
	FindAll() []Todo
	FindByActivityID(id int) ([]Todo, error)
	DeleteByID(id int) error
	UpdateByID(todo Todo) (Todo, error)
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

func (r *todoRepository) FindAll() []Todo {
	var todos []Todo

	r.db.Find(&todos)

	return todos
}

func (r *todoRepository) FindByActivityID(id int) ([]Todo, error) {
	var todos []Todo

	if err := r.db.Where("activity_group_id = ?", id).Find(&todos).Error; err != nil {
		return todos, err
	}

	return todos, nil
}

func (r *todoRepository) DeleteByID(id int) error {
	var todo Todo

	if err := r.db.Where("id = ?", id).Delete(&todo).Error; err != nil {
		return err
	}

	return nil
}

func (r *todoRepository) UpdateByID(todo Todo) (Todo, error) {
	if err := r.db.Save(&todo).Error; err != nil {
		return todo, err
	}

	return todo, nil
}
