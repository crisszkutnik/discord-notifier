package main

import (
	"github.com/crisszkutnik/discord-notifier/internal/env"
	"github.com/crisszkutnik/discord-notifier/internal/http"
	"github.com/gofiber/fiber/v3"
)

func main() {
	env.LoadEnv()

	app := fiber.New()

	http.RegisterRouter(app)

	app.Listen(*env.PORT)
}
