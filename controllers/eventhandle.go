package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jelmerdereus/gowebtemplate/datastore"
	"github.com/jelmerdereus/gowebtemplate/models"
)

//EventHandle contains the orm layer for User objects
type EventHandle struct {
	Handle
}

//NewEventHandle takes an ORM and returns an NewEventHandle
func NewEventHandle(orm *datastore.DBORM) (*EventHandle, error) {
	if orm == nil {
		return nil, errors.New("No ORM provided")
	}
	orm.AutoMigrate(&models.Event{})
	return &EventHandle{Handle: Handle{DB: orm}}, nil
}

//Create creates an Event
func (e *EventHandle) Create(c *gin.Context) {
	action := "Event.Create"
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(action, "InputError", err))
		return
	}
	user, err := e.DB.AddEvent(event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(action, "DBError", err))
		return
	}

	c.JSON(http.StatusCreated, SuccessResponse(action, user))
}

//GetAll gets all events
func (e *EventHandle) GetAll(c *gin.Context) {
	action := "Event.GetAll"

	events, err := e.DB.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(action, "DBerror", err))
		return
	}
	c.JSON(http.StatusAccepted, SuccessResponse(action, events))
}

//GetByID gets an event by its ID
func (e *EventHandle) GetByID(c *gin.Context) {
	action := "Event.GetByID"

	idstring := c.Param("id")
	if idstring == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse(action, "InputError", errors.New("param id is required")))
		return
	}
	id, err := strconv.Atoi(idstring)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(action, "InputError", err))
		return
	}
	event, err := e.DB.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(action, "DBerror", err))
		return
	}
	c.JSON(http.StatusAccepted, SuccessResponse(action, event))
}
