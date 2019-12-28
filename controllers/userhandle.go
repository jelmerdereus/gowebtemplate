package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jelmerdereus/goweb3/datastore"
	"github.com/jelmerdereus/goweb3/models"
)

//UserHandle contains the orm layer for User objects
type UserHandle struct {
	Handle
}

//NewUserHandle takes an ORM and returns a UserHandle
func NewUserHandle(orm *datastore.DBORM) (*UserHandle, error) {
	if orm == nil {
		return nil, errors.New("No ORM provided")
	}
	orm.AutoMigrate(&models.User{})
	return &UserHandle{Handle: Handle{DB: orm}}, nil
}

// GetAll returns an array of all users
func (u *UserHandle) GetAll(c *gin.Context) {
	users, err := u.DB.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse("Users.GetAll", "DBError", err))
		return
	}
	c.JSON(http.StatusAccepted, SuccessResponse("User.GetAll", users))
}

// GetByAlias returns a user with a certain alias
func (u *UserHandle) GetByAlias(c *gin.Context) {
	user, err := u.DB.GetUserByAlias(c.Param("alias"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse("User.GetByAlias", "DBError", err))
		return
	}
	c.JSON(http.StatusAccepted, SuccessResponse("User.GetByAlias", user))
}

// GetByID returns a user with a certain id
func (u *UserHandle) GetByID(c *gin.Context) {
	idstring, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, ErrorResponse("User.GetByID", "InputError", errors.New("parameter id is required")))
		return
	}
	userid, err := strconv.Atoi(idstring)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse("User.GetByID", "InputError", errors.New("parameter id is invalid")))
		return
	}
	user, err := u.DB.GetUserByID(userid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse("User.GetByID", "DBError", err))
		return
	}
	c.JSON(http.StatusAccepted, SuccessResponse("User.GetByID", user))
}

// Create adds a user and returns it
func (u *UserHandle) Create(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse("User.Create", "InputError", err))
		return
	}
	user, err := u.DB.AddUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse("User.Create", "DBError", err))
		return
	}

	c.JSON(http.StatusCreated, SuccessResponse("User.Create", user))
}

//Update updates the properties of a user
func (u *UserHandle) Update(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse("User.Update", "InputError", errors.New("parameter id not provided")))
		return
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse("User.Update", "InputError", errors.New("parameter id invalid")))
		return
	}

	var user models.User

	// verify that the body can be parsed
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse("User.Update", "InputError", err))
		return
	}

	// find the user
	if _, err := u.DB.GetUserByID(id); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse("User.Update", "DBError", err))
		return
	}

	// update the user
	if err := u.DB.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse("User.Update", "DBError", err))
		return
	}

	c.JSON(http.StatusAccepted, SuccessResponse("User.Update", user))
}

// Delete deletes a user
func (u *UserHandle) Delete(c *gin.Context) {
	// input validation
	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse("User.Update", "InputError", errors.New("parameter id not provided")))
		return
	}

	var user models.User

	userid, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse("User.Delete", "InputError", errors.New("parameter id invalid")))
		return
	}

	// find the user
	user, err = u.DB.GetUserByID(userid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse("User.Delete", "DBError", err))
		return
	}

	// delete it
	if err := u.DB.DeleteUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse("User.Delete", "DBError", err))
		return
	}
	var deleted = time.Now().UTC()
	user.DeletedAt = &deleted

	c.JSON(http.StatusAccepted, SuccessResponse("User.Delete", user))
}
