package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// test the controller for the / ROOT path
func TestGetRoot(t *testing.T) {
	router := gin.New()
	router.GET("/", GetRoot)

	recorder := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(recorder, req)

	body, _ := ioutil.ReadAll(recorder.Body)
	data := make(map[string]interface{})
	_ = json.Unmarshal(body, &data)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "success", data["status"])
	assert.Equal(t, "GetRoot", data["action"])
	assert.Equal(t, "Hello, World!", data["result"])
}
