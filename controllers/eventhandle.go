package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jelmerdereus/gowebtemplate/datastore"
	"github.com/jelmerdereus/gowebtemplate/models"
)

// EventAPI is the interface for the Event REST methods
type EventAPI interface {
	RESTAPI
}

//EventHandle contains the orm layer for User objects
type EventHandle struct {
	Repo datastore.EventRepo
}

//NewEventAPI takes an ORM and returns an NewEventHandle
func NewEventAPI(repo datastore.EventRepo) EventAPI {
	return &EventHandle{Repo: repo}
}

//GetAll gets all events
func (e *EventHandle) GetAll(c *gin.Context) {
	action := "Event.GetAll"

	events, err := e.Repo.GetAllEvents()
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
	event, err := e.Repo.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(action, "DBerror", err))
		return
	}
	c.JSON(http.StatusAccepted, SuccessResponse(action, event))
}

//Create creates an Event
func (e *EventHandle) Create(c *gin.Context) {
	action := "Event.Create"

	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(action, "InputError", err))
		return
	}
	event, err := e.Repo.AddEvent(event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(action, "DBError", err))
		return
	}

	c.JSON(http.StatusCreated, SuccessResponse(action, event))
}

//Update updates an Event and returns it
func (e *EventHandle) Update(c *gin.Context) {
	action := "Event.Update"

	// id parameter validation
	idString := c.Param("id")
	if idString == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse(action, "InputError", errors.New("param id is required")))
		return
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(action, "InputError", errors.New("param id is invalid")))
		return
	}

	// parse the body
	var updatedEvent models.Event
	if err = c.ShouldBindJSON(&updatedEvent); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(action, "InputError", err))
		return
	}

	// find the event
	event, err := e.Repo.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(action, "DBError", err))
		return
	}

	// update it
	if err = e.Repo.UpdateEvent(event); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(action, "DBError", err))
		return
	}

	c.JSON(http.StatusAccepted, SuccessResponse(action, event))
}

//Delete deletes an event and returns the object
func (e *EventHandle) Delete(c *gin.Context) {
	action := "Event.Delete"

	// id parameter validation
	idString := c.Param("id")
	if idString == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse(action, "InputError", errors.New("param id is required")))
		return
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(action, "InputError", errors.New("param id is invalid")))
		return
	}

	// find the event
	event, err := e.Repo.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(action, "DBError", err))
		return
	}

	// delete it
	err = e.Repo.DeleteEvent(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(action, "DBError", err))
		return
	}

	var deleted = time.Now().UTC()
	event.DeletedAt = &deleted

	c.JSON(http.StatusAccepted, SuccessResponse(action, event))
}
