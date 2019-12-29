package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jelmerdereus/gowebtemplate/datastore"
	"net/http"
)

//RESTAPI is an interface for implementation by REST controllers
type RESTAPI interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

//Handle contains the orm layer and is embedded by the implementations of RESTAPI
type Handle struct {
	DB datastore.DBLayer
}

// GetRoot handles requests to the root url
func GetRoot(c *gin.Context) {
	c.JSON(http.StatusAccepted, SuccessResponse("GetRoot", "Hello, World!"))
}

//SuccessResponse returns a map that can be given as an argument to c.JSON
func SuccessResponse(action string, result interface{}) gin.H {
	return gin.H{
		"action": action,
		"status": "success",
		"result": result,
	}
}

//ErrorResponse returns a map that can be given as an argument to c.JSON
func ErrorResponse(action, errorType string, err error) gin.H {
	return gin.H{
		"action": action,
		"status": "error",
		"error": gin.H{
			"type":    errorType,
			"message": err.Error(),
		},
	}
}
