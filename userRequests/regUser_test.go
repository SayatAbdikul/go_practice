package userRequests_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SayatAbdikul/go_practice/server"
	"github.com/SayatAbdikul/go_practice/userRequests"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegUser(t *testing.T) {
	server.Connect()
	router := gin.Default()
	router.POST("/reg_user", userRequests.RegUser)
	user := userRequests.User{Name: "John", Login: "some_user", Password: "test_password", Age: 90}
	load, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}
	request, err := http.NewRequest(http.MethodPost, "/reg_user", bytes.NewBuffer(load))
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-type", "application/json")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code)
	expected := ``
	assert.Equal(t, expected, recorder.Body.String())
}
