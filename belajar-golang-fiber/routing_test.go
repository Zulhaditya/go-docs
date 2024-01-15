package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var app = fiber.New()

func TestRoutingHelloFiber(t *testing.T) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello Fiber!")
	})

	request := httptest.NewRequest("GET", "/", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Fiber!", string(bytes))
}

func TestCtx(t *testing.T) {
	app.Get("/hello", func(ctx *fiber.Ctx) error {
		name := ctx.Query("name", "Guest")
		return ctx.SendString("Hello " + name)
	})

	request := httptest.NewRequest("GET", "/hello?name=Ackxle", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Ackxle", string(bytes))

	request = httptest.NewRequest("GET", "/hello", nil)
	response, err = app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err = io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Guest", string(bytes))

}

func TestHTTPRequest(t *testing.T) {
	app.Get("/request", func(ctx *fiber.Ctx) error {
		first := ctx.Get("firstname")
		last := ctx.Cookies("lastname")
		return ctx.SendString("Hello " + first + " " + last)
	})

	request := httptest.NewRequest("GET", "/request", nil)
	request.Header.Set("firstname", "Ackxle")
	request.AddCookie(&http.Cookie{Name: "lastname", Value: "Inayah"})
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Ackxle Inayah", string(bytes))
}

func TestRouteParameter(t *testing.T) {
	app.Get("/users/:userId/orders/:orderId", func(ctx *fiber.Ctx) error {
		userId := ctx.Params("userId")
		orderId := ctx.Params("orderId")
		return ctx.SendString("Get order " + orderId + " from user " + userId)
	})

	request := httptest.NewRequest("GET", "/users/1/orders/2", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Get order 2 from user 1", string(bytes))
}

func TestFormRequest(t *testing.T) {
	app.Post("/hello", func(ctx *fiber.Ctx) error {
		name := ctx.FormValue("name")
		return ctx.SendString("Hello " + name)
	})

	body := strings.NewReader("name=Ackxle")
	request := httptest.NewRequest("POST", "/hello", body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Ackxle", string(bytes))
}

// Implementasi multipart-form

//go:embed source/contoh.txt
var contohFile []byte

func TestFormUpload(t *testing.T) {
	app.Post("/upload", func(ctx *fiber.Ctx) error {
		file, err := ctx.FormFile("file")
		if err != nil {
			return err
		}

		err = ctx.SaveFile(file, "./target/"+file.Filename)
		if err != nil {
			return err
		}

		return ctx.SendString("Upload success!")
	})

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	file, _ := writer.CreateFormFile("file", "contoh.txt")
	file.Write(contohFile)
	writer.Close()

	request := httptest.NewRequest("POST", "/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	response, err := app.Test(request)
	assert.Nil(t, err)
	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Upload success!", string(bytes))

}

// implementasi request body
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestRequestBody(t *testing.T) {
	app.Post("login", func(ctx *fiber.Ctx) error {
		body := ctx.Body()
		request := new(LoginRequest)
		err := json.Unmarshal(body, request)
		if err != nil {
			return err
		}

		return ctx.SendString("Login Success " + request.Username)
	})

	body := strings.NewReader(`{"username":"Ackxle","password":"secret"}`)
	request := httptest.NewRequest("POST", "/login", body)
	request.Header.Set("Content-Type", "application/json")
	response, err := app.Test(request)
	assert.Nil(t, err)
	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Login Success Ackxle", string(bytes))
}