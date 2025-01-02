package main

import "github.com/gofiber/fiber/v2"

func main() {
	app:=fiber.New()
	app.Static("/","./Frontend")
	println("http://localhost:5430")
	app.Listen(":5430")
}