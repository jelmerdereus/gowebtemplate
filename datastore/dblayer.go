package datastore

import (
	"github.com/jelmerdereus/gowebtemplate/models"
)

//DBLayer is the abstraction layer for the database
type DBLayer interface {
	// Users
	GetAllUsers() ([]models.User, error)
	GetUserByAlias(alias string) (models.User, error)
	GetUserByID(id int) (models.User, error)
	AddUser(models.User) (models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(models.User) error

	// Events
	GetAllEvents() ([]models.Event, error)
	GetEventByID(id int) (models.Event, error)
	AddEvent(models.Event) (models.Event, error)
	UpdateEvent(models.Event) error
	DeleteEvent(*models.Event) error
}
