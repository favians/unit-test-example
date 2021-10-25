package controllers_test

import (
	"fmt"
	"mvc/config"
	"mvc/models"
	"mvc/routes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gavv/httpexpect"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func PrepareTestData() error {
	config.InitDB()

	if os.Getenv("ENV") == "test" {
		config.DB.Exec("TRUNCATE TABLE users;")
	} else {
		return fmt.Errorf("invalid Environment")
	}

	config.DB.Create(models.Users{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:     "unit-test",
		Email:    "unit-test",
		Password: "unit-test",
		Token:    "unit-test",
	})

	return nil
}

func InitEcho() *echo.Echo {

	e := routes.New()

	return e
}

func TestGetUserByIDControllers(t *testing.T) {

	PrepareTestData()

	handler := InitEcho()

	// run server using httptest
	server := httptest.NewServer(handler)

	defer server.Close()

	// create httpexpect for Mock instance
	e := httpexpect.New(t, server.URL)

	//==================GET JWT TOKEN FOR ADD IN HEADER REQUEST===================
	data := map[string]interface{}{
		"email":    "unit-test",
		"password": "unit-test",
	}
	// get token
	obj := e.POST("/login").
		WithJSON(data).
		Expect().
		Status(http.StatusOK).JSON().Object()

	token := obj.Value("users").Object().Value("token").String().Raw()

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})

	t.Run("Expected Found the user find by ID", func(t *testing.T) {
		auth.GET("/users").WithQuery("id", "1").
			Expect().
			Status(http.StatusOK).JSON().Object().
			Value("users").Object().Value("ID").Equal(1)
	})

	t.Run("Expected user find by ID NOT FOUND", func(t *testing.T) {
		auth.GET("/users").WithQuery("id", "999").
			Expect().
			Status(http.StatusNotFound)
	})
}

func TestGetUserControllers(t *testing.T) {

	PrepareTestData()

	handler := InitEcho()

	// run server using httptest
	server := httptest.NewServer(handler)

	defer server.Close()

	// create httpexpect for Mock instance
	e := httpexpect.New(t, server.URL)

	//==================GET JWT TOKEN FOR ADD IN HEADER REQUEST===================
	data := map[string]interface{}{
		"email":    "unit-test",
		"password": "unit-test",
	}
	// get token
	obj := e.POST("/login").
		WithJSON(data).
		Expect().
		Status(http.StatusOK).JSON().Object()

	token := obj.Value("users").Object().Value("token").String().Raw()

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})

	t.Run("Expected Find ALL get user", func(t *testing.T) {
		auth.GET("/users/all").
			Expect().
			Status(http.StatusOK).JSON().Object().Value("users").Array().
			Element(0).Object().Value("ID").Equal(1)
	})
}

func TestInsertUserControllers(t *testing.T) {

	PrepareTestData()

	handler := InitEcho()

	// run server using httptest
	server := httptest.NewServer(handler)

	defer server.Close()

	// create httpexpect for Mock instance
	e := httpexpect.New(t, server.URL)

	//==================GET JWT TOKEN FOR ADD IN HEADER REQUEST===================
	data := map[string]interface{}{
		"email":    "unit-test",
		"password": "unit-test",
	}
	// get token
	obj := e.POST("/login").
		WithJSON(data).
		Expect().
		Status(http.StatusOK).JSON().Object()

	token := obj.Value("users").Object().Value("token").String().Raw()

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})

	t.Run("Expected Insert user, then call get one By ID and found that user", func(t *testing.T) {
		dataForInsert := map[string]interface{}{
			"name":     "user-1",
			"email":    "user-1@gmail.com",
			"password": "12345678",
		}

		auth.POST("/users").WithJSON(dataForInsert).
			Expect().
			Status(http.StatusOK)

		auth.GET("/users").WithQuery("id", 2).
			Expect().
			Status(http.StatusOK).JSON().Object().
			Value("users").Object().Value("name").Equal("user-1")
	})

	t.Run("Expected Insert user, then call get one By ID But NOT found that user", func(t *testing.T) {
		dataForInsert := map[string]interface{}{
			"name":     "user-1",
			"email":    "user-1@gmail.com",
			"password": "12345678",
		}

		auth.POST("/users").WithJSON(dataForInsert).
			Expect().
			Status(http.StatusOK)

		auth.GET("/users").WithQuery("id", 999).
			Expect().
			Status(http.StatusNotFound)
	})

}
