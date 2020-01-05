package datastore

import (
	"errors"

	"github.com/jelmerdereus/gowebtemplate/models"
	"github.com/jinzhu/gorm"
)

// UserStore is an ORM layer that satisfies the UserRepo interface
type UserStore struct {
	DBORM
}

// NewUserRepo is a constructor
func NewUserRepo(orm *DBORM) (UserRepo, error) {
	if orm == nil {
		return nil, errors.New("No ORM provided")
	}
	orm.AutoMigrate(&models.User{})

	store := UserStore{}
	store.DB = orm.DB

	return &store, nil
}

// GetAllUsers returns a list of User objects
func (store *UserStore) GetAllUsers() (users []models.User, err error) {
	return users, store.Find(&users).Error
}

//GetUserByAlias returns the User object of the user with the given alias
func (store *UserStore) GetUserByAlias(alias string) (user models.User, err error) {
	return user, store.First(&user, &models.User{Alias: alias}).Error
}

//GetUserByID returns the User object of the user with the given ID
func (store *UserStore) GetUserByID(id int) (user models.User, err error) {
	return user, store.First(&user, &models.User{Model: gorm.Model{ID: uint(id)}}).Error
}

// AddUser adds a User object to the database and returns the user
func (store *UserStore) AddUser(newUser models.User) (models.User, error) {
	return newUser, store.Create(&newUser).Error
}

// UpdateUser updates a User object and returns it
func (store *UserStore) UpdateUser(user *models.User) error {
	return store.Save(user).Error
}

// DeleteUser deletes a User object and returns the object with DeletedAt timestamp
func (store *UserStore) DeleteUser(user models.User) error {
	return store.Delete(&user).Error
}
