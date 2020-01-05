package datastore

import (
	"errors"

	"github.com/jelmerdereus/gowebtemplate/models"
	"github.com/jinzhu/gorm"
)

// EventStore is an ORM layer that satisfies the DBLayer interface
type EventStore struct {
	DBORM
}

// NewEventRepo is a constructor for DBORM
func NewEventRepo(orm *DBORM) (EventRepo, error) {
	if orm == nil {
		return nil, errors.New("No ORM provided")
	}
	orm.AutoMigrate(&models.Event{})

	store := EventStore{}
	store.DB = orm.DB

	return &store, nil
}

// GetAllEvents returns all events
func (store *EventStore) GetAllEvents() (events []models.Event, err error) {
	return events, store.Find(&events).Error
}

// GetEventByID returns an event or an error
func (store *EventStore) GetEventByID(id int) (event models.Event, err error) {
	return event, store.First(&event, &models.Event{Model: gorm.Model{ID: uint(id)}}).Error
}

// AddEvent adds an event and returns it
func (store *EventStore) AddEvent(event models.Event) (models.Event, error) {
	return event, store.Create(&event).Error
}

// UpdateEvent updates an Event object
func (store *EventStore) UpdateEvent(event models.Event) error {
	return store.Save(event).Error
}

//DeleteEvent deletes an event object and returns it
func (store *EventStore) DeleteEvent(event *models.Event) error {
	return store.Delete(&event).Error
}
