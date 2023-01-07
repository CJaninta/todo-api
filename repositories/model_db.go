package repositories

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
}

func (b *BaseModel) BeforeCreate(db *gorm.DB) (err error) {
	b.CreatedAt = time.Unix(time.Now().Unix(), 0)
	b.UpdatedAt = time.Unix(time.Now().Unix(), 0)
	return
}

func (b *BaseModel) BeforeUpdate(db *gorm.DB) (err error) {
	b.UpdatedAt = time.Unix(time.Now().Unix(), 0)
	return
}

type UserProjectRepository interface {
	AddUserInProject(int, int) error
	DeleteUserInProject(int, int) error
}

type userProjectRepositoryDB struct {
	db *gorm.DB
}

func NewUserProjectRepositoryDB(db *gorm.DB) UserProjectRepository {
	return userProjectRepositoryDB{db: db}
}

func (r userProjectRepositoryDB) AddUserInProject(userId int, projectId int) error {

	user, project, err := FindUserAndProject(userId, projectId, r.db)
	if err != nil {
		fmt.Println(err)
		return err
	}

	//TODO: should optimize
	findProject := Project{}
	_ = r.db.Model(&user).Where("id=?", projectId).Association("Projects").Find(&findProject)
	if findProject.ID != 0 {
		fmt.Println(err)
		return errors.New("User already have this project")
	}

	projects := []Project{
		*project,
	}
	err = r.db.Model(&user).Where("id=?", projectId).Association("Projects").Append(&projects)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (r userProjectRepositoryDB) DeleteUserInProject(userId int, projectId int) error {

	user, project, err := FindUserAndProject(userId, projectId, r.db)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = r.db.Model(&user).Association("Projects").Delete(&project)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

//TODO: should optimize
func FindUserAndProject(userId int, projectId int, db *gorm.DB) (*User, *Project, error) {
	project := Project{}
	tx := db.Where("id=?", projectId).First(&project)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, nil, tx.Error
	}

	user := User{}
	tx = db.Where("id=?", userId).First(&user)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, nil, tx.Error
	}

	return &user, &project, nil
}
