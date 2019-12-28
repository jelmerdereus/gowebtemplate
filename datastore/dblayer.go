package datastore

import (
	"github.com/jelmerdereus/goweb3/models"
)

//DBLayer is the abstraction layer for the database
type DBLayer interface {
	GetAllUsers() (users []models.User, err error)
	GetUserByAlias(alias string) (user models.User, err error)
	GetUserByID(id int) (user models.User, err error)
	AddUser(models.User) (user models.User, err error)
	UpdateUser(user *models.User) error
	DeleteUser(user models.User) error
}
