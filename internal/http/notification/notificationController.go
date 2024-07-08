package notification

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v3"
)

const ChannelIdParamName = "channelId"

type NotificationController struct {
	notificationService *NotificationService
}

func NewNotificationController(notificationService *NotificationService) *NotificationController {
	return &NotificationController{
		notificationService: notificationService,
	}
}

func (c *NotificationController) SendNotification(ctx fiber.Ctx) error {
	channelId := ctx.Params(ChannelIdParamName)

	p := new(discordgo.MessageSend)

	if err := ctx.Bind().Body(p); err != nil {
		return err
	}

	err := c.notificationService.SendNotification(channelId, p)

	if err != nil {
		return err
	}

	return ctx.SendStatus(200)
}
