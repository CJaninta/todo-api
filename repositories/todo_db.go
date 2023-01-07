package repositories

import (
	"fmt"

	"gorm.io/gorm"
)

type todoRepositoryDB struct {
	db *gorm.DB
}

func NewTodoRepositoryDB(db *gorm.DB) TodoRepository {
	return todoRepositoryDB{db: db}
}

func (r todoRepositoryDB) GetAll() ([]Todo, error) {
	todos := []Todo{}
	tx := r.db.Order("id").Find(&todos)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return todos, nil
}

func (r todoRepositoryDB) GetById(id int) (*Todo, error) {
	todo := Todo{}
	tx := r.db.Where("id=?", id).First(&todo)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}
	return &todo, nil
}

func (r todoRepositoryDB) GetByProjectId(projectId int) ([]Todo, error) {
	todos := []Todo{}
	tx := r.db.Where("project_id=?", projectId).Find(&todos)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}
	return todos, nil
}

func (r todoRepositoryDB) Create(todo Todo) (*Todo, error) {

	newTodo := Todo{
		Title:     todo.Title,
		Detail:    todo.Detail,
		Step:      todo.Step,
		UserID:    todo.UserID,
		ProjectID: todo.ProjectID,
	}

	tx := r.db.Create(&newTodo)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}
	return &newTodo, nil
}

func (r todoRepositoryDB) UpdateStep(id int, step string) error {

	todo, err := r.GetById(id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	tx := r.db.Model(&todo).Select("Step").Updates(Todo{Step: step})
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}

	return nil
}
