package datastore

import (
	"github.com/jelmerdereus/gowebtemplate/models"
	"github.com/jinzhu/gorm"
)

// DBORM is an ORM layer that satisfies the DBLayer interface
type DBORM struct {
	*gorm.DB
}

// NewORM is a constructor for DBORM
func NewORM(dbname, con string) (*DBORM, error) {
	db, err := gorm.Open(dbname, con)
	return &DBORM{DB: db}, err
}

//UserRepo is the abstraction for user related database operations
type UserRepo interface {
	GetAllUsers() ([]models.User, error)
	GetUserByAlias(alias string) (models.User, error)
	GetUserByID(id int) (models.User, error)
	AddUser(models.User) (models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(models.User) error
}

//EventRepo is the abstraction for event related database operations
type EventRepo interface {
	GetAllEvents() ([]models.Event, error)
	GetEventByID(id int) (models.Event, error)
	AddEvent(models.Event) (models.Event, error)
	UpdateEvent(models.Event) error
	DeleteEvent(*models.Event) error
}
