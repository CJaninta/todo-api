package repositories

import (
	"fmt"

	"gorm.io/gorm"
)

type projectRepositoryDB struct {
	db *gorm.DB
}

func NewProjectRepositoryDB(db *gorm.DB) ProjectRepository {
	return projectRepositoryDB{db: db}
}

func (r projectRepositoryDB) GetAll() ([]Project, error) {
	projects := []Project{}
	tx := r.db.Order("id").Find(&projects)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return projects, nil
}

func (r projectRepositoryDB) GetById(id uint) (*Project, error) {
	project := Project{}
	tx := r.db.Where("id=?", id).First(&project)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return &project, nil
}

func (r projectRepositoryDB) GetByName(name string) (*Project, error) {
	project := Project{}
	tx := r.db.Where("name=?", name).First(&project)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return &project, nil
}

func (r projectRepositoryDB) Create(project Project) (*Project, error) {
	newProject := Project{
		Name:   project.Name,
		Detail: project.Detail,
	}

	tx := r.db.Create(&newProject)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return &newProject, nil
}

func (r projectRepositoryDB) GetProjectByUser(userId int) ([]Project, error) {
	projects := []Project{}
	user := User{}

	tx := r.db.Where("id=?", userId).First(&user)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	err := r.db.Model(&user).Association("Projects").Find(&projects)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return projects, err
}
