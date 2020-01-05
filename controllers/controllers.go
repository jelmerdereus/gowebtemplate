package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//EveryBody is the body that is always returned
type EveryBody struct {
	Action string `json:"action"`
	Status string `json:"status"`
}

//SuccessBody is the body that is returned on success
type SuccessBody struct {
	EveryBody
	Result interface{} `json:"result"`
}

//ErrorDetails are details about errors
type ErrorDetails struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

//ErrorBody is the body that is returned on errors
type ErrorBody struct {
	EveryBody
	Error ErrorDetails `json:"error"`
}

//RESTAPI is an interface for implementation by REST controllers
type RESTAPI interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

// GetRoot handles requests to the root url
func GetRoot(c *gin.Context) {
	c.JSON(http.StatusAccepted, SuccessResponse("GetRoot", "Hello, World!"))
}

//SuccessResponse returns a map that can be given as an argument to c.JSON
func SuccessResponse(action string, result interface{}) SuccessBody {
	body := SuccessBody{}
	body.Action = action
	body.Status = "success"
	body.Result = result

	return body
}

//ErrorResponse returns a map that can be given as an argument to c.JSON
func ErrorResponse(action, errorType string, err error) ErrorBody {
	body := ErrorBody{}
	body.Action = action
	body.Status = "error"
	body.Error.Type = errorType
	body.Error.Message = err.Error()

	return body
}
