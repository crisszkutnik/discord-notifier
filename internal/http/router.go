package http

import (
	"github.com/crisszkutnik/discord-notifier/internal/discord"
	"github.com/crisszkutnik/discord-notifier/internal/http/notification"
	"github.com/gofiber/fiber/v3"
)

func RegisterRouter(app *fiber.App) {
	discordService := discord.NewDiscordService()
	notificationService := notification.NewNotificationService(discordService)
	notificationController := notification.NewNotificationController(notificationService)

	notificationGroup := app.Group("/notification")

	notificationGroup.Post("/discord/:"+notification.ChannelIdParamName, notificationController.SendNotification)
}
