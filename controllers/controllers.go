package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jelmerdereus/goweb3/datastore"
	"net/http"
)

//Handle contains the orm layer
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
