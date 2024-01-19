package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	})

	app.Use(func(ctx *fiber.Ctx) error {
		fmt.Println("i'm middleware before processing request")
		err := ctx.Next()
		fmt.Println("i'm middleware after processing request")
		return err
	})

	// menjalankan middleware pada route api tertentu
	app.Use("/api", func(ctx *fiber.Ctx) error {
		fmt.Println("i'm middleware before processing request")
		err := ctx.Next()
		fmt.Println("i'm middleware after processing request")
		return err
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World")
	})

	err := app.Listen("localhost:3000")
	if err != nil {
		panic(err)
	}
}
