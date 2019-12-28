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

// GetAllUsers returns a list of User objects
func (orm *DBORM) GetAllUsers() (users []models.User, err error) {
	return users, orm.Find(&users).Error
}

//GetUserByAlias returns the User object of the user with the given alias
func (orm *DBORM) GetUserByAlias(alias string) (user models.User, err error) {
	return user, orm.First(&user, &models.User{Alias: alias}).Error
}

//GetUserByID returns the User object of the user with the given ID
func (orm *DBORM) GetUserByID(id int) (user models.User, err error) {
	return user, orm.First(&user, &models.User{Model: gorm.Model{ID: uint(id)}}).Error
}

// AddUser adds a User object to the database and returns the user
func (orm *DBORM) AddUser(newUser models.User) (models.User, error) {
	return newUser, orm.Create(&newUser).Error
}

// UpdateUser updates a User object and returns it
func (orm *DBORM) UpdateUser(user *models.User) error {
	return orm.Save(user).Error
}

// DeleteUser deletes a User object and returns the object with DeletedAt timestamp
func (orm *DBORM) DeleteUser(user models.User) error {
	return orm.Delete(&user).Error
}

// GetAllEvents returns all events
func (orm *DBORM) GetAllEvents() (events []models.Event, err error) {
	return events, orm.Find(&events).Error
}

// GetEventByID returns an event or an error
func (orm *DBORM) GetEventByID(id int) (event models.Event, err error) {
	return event, orm.First(&event, &models.Event{Model: gorm.Model{ID: uint(id)}}).Error
}

// AddEvent adds an event and returns it
func (orm *DBORM) AddEvent(event models.Event) (models.Event, error) {
	return event, orm.Create(&event).Error
}
