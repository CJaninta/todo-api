package repositories

import (
	"fmt"

	"gorm.io/gorm"
)

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) UserRepository {
	return userRepositoryDB{db: db}
}

func (r userRepositoryDB) GetAll() ([]User, error) {
	users := []User{}
	tx := r.db.Order("id").Find(&users)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return users, nil
}

func (r userRepositoryDB) GetById(id uint) (*User, error) {
	user := User{}
	tx := r.db.Where("id=?", id).First(&user)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return &user, nil
}

func (r userRepositoryDB) GetByEmail(email string) (*User, error) {
	user := User{}
	tx := r.db.Where("email=?", email).First(&user)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return &user, nil
}

func (r userRepositoryDB) Create(user User) (*User, error) {
	newUser := User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}

	tx := r.db.Create(&newUser)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return &newUser, nil
}
