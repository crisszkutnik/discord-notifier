package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/crisszkutnik/discord-notifier/internal/env"
	"github.com/gofiber/fiber/v3"
)

type DiscordService struct {
	client *discordgo.Session
}

func NewDiscordService() *DiscordService {
	client, err := discordgo.New(*env.DISCORD_TOKEN)

	if err != nil {
		log.Fatal("Could not initialize Discord client")
	}

	return &DiscordService{
		client: client,
	}
}

func (s *DiscordService) SendComplexMessage(channelId string, payload *discordgo.MessageSend) error {
	_, err := s.client.ChannelMessageSendComplex(channelId, payload)

	if err != nil {
		log.Print("Error sending Discord message. ", err)

		if isBadRequest(err) {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return fiber.ErrInternalServerError
	}

	return nil
}
