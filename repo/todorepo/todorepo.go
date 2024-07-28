package todorepo

import (
	todosModel "anna/models/todos"
	"anna/repo/userrepo"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID     int    `gorm:"id"`
	Task   string `gorm:"task"`
	Status bool   `gorm:"status"`
	Email  string `gorm:"email"`
}

type TodoRepo struct {
	db *gorm.DB
}

func NewTodoRepo(db *gorm.DB) *TodoRepo {
	return &TodoRepo{db: db}
}

func (todoRepo *TodoRepo) CreateTodo(newTodo *todosModel.Todo, userEmail string) (*Todo, error) {
	var user userrepo.User
	if err := todoRepo.db.Select("ID").Where("email = ?", userEmail).Find(&user).Error; err != nil {
		return nil, err
	}
	var todo = &Todo{
		Task:   newTodo.Task,
		Email:  userEmail,
		Status: newTodo.Status,
	}
	if err := todoRepo.db.Create(todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func (todoRepo *TodoRepo) GetTodos(userEmail string) ([]Todo, error) {
	todosList := []Todo{}

	if err := todoRepo.db.Where("email = ?", userEmail).Find(&todosList).Error; err != nil {
		return nil, err
	}
	return todosList, nil
}

func (todoRepo *TodoRepo) UpdateTodo(todoId int, updateTodo Todo) (*Todo, error) {
	if err := todoRepo.db.Model(&Todo{}).Where("ID = ?", todoId).Update("status", updateTodo.Status).Error; err != nil {
		return nil, err
	}
	return todoRepo.GetTodoById(todoId)
}

func (todoRepo *TodoRepo) GetTodoById(todoId int) (*Todo, error) {
	var todo Todo
	if err := todoRepo.db.Where("ID = ?", todoId).First(&todo).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (todoRepo *TodoRepo) DeleteTodo(todoId int) error {
	if err := todoRepo.db.Delete(&todosModel.Todo{}, todoId).Error; err != nil {
		return err
	}
	return nil
}
