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

//NewUserHandle takes an ORM and returns a UserHandle
func NewUserHandle(orm *datastore.DBORM) (*Handle, error) {
	if orm == nil {
		return nil, errors.New("No ORM provided")
	}
	orm.AutoMigrate(&models.User{})
	return &Handle{DB: orm}, nil
}

// GetAll returns an array of all users
func (h *Handle) GetAll(c *gin.Context) {
	action := "Users.GetAll"
	users, err := h.DB.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(action, "DBError", err))
		return
	}
	c.JSON(http.StatusAccepted, SuccessResponse(action, users))
}

// GetByAlias returns a user with a certain alias
func (h *Handle) GetByAlias(c *gin.Context) {
	action := "User.GetByAlias"
	user, err := h.DB.GetUserByAlias(c.Param("alias"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(action, "DBError", err))
		return
	}
	c.JSON(http.StatusAccepted, SuccessResponse(action, user))
}

// GetByID returns a user with a certain id
func (h *Handle) GetByID(c *gin.Context) {
	action := "User.GetByID"
	idstring, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, ErrorResponse(action, "InputError", errors.New("parameter id is required")))
		return
	}
	userid, err := strconv.Atoi(idstring)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(action, "InputError", errors.New("parameter id is invalid")))
		return
	}
	user, err := h.DB.GetUserByID(userid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(action, "DBError", err))
		return
	}
	c.JSON(http.StatusAccepted, SuccessResponse(action, user))
}

// Create adds a user and returns it
func (h *Handle) Create(c *gin.Context) {
	action := "User.Create"
	var user models.User

	// parse the object
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(action, "InputError", err))
		return
	}

	// add the user to the database
	user, err := h.DB.AddUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(action, "DBError", err))
		return
	}

	// transmit the response
	c.JSON(http.StatusCreated, SuccessResponse(action, user))

	// add an event to the database
	h.DB.AddEvent(models.Event{
		UserID:  user.ID,
		Action:  action,
		Type:    "DB",
		Message: "success",
	})
}

//Update updates the properties of a user
func (h *Handle) Update(c *gin.Context) {
	action := "User.Update"

	// input validation
	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse(action, "InputError", errors.New("parameter id not provided")))
		return
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(action, "InputError", errors.New("parameter id invalid")))
		return
	}

	var user models.User

	// verify that the body can be parsed
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(action, "InputError", err))
		return
	}

	// find the user
	if _, err := h.DB.GetUserByID(id); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(action, "DBError", err))
		return
	}

	// update the user
	if err := h.DB.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(action, "DBError", err))
		return
	}

	c.JSON(http.StatusAccepted, SuccessResponse(action, user))
}

// Delete deletes a user
func (h *Handle) Delete(c *gin.Context) {
	action := "User.Delete"

	// input validation
	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse(action, "InputError", errors.New("parameter id not provided")))
		return
	}

	var user models.User

	userid, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(action, "InputError", errors.New("parameter id invalid")))
		return
	}

	// find the user
	user, err = h.DB.GetUserByID(userid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(action, "DBError", err))
		return
	}

	// delete it
	if err := h.DB.DeleteUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(action, "DBError", err))
		return
	}
	var deleted = time.Now().UTC()
	user.DeletedAt = &deleted

	c.JSON(http.StatusAccepted, SuccessResponse(action, user))
}
