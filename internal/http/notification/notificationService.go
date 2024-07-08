package notification

import (
	"github.com/bwmarrin/discordgo"
	"github.com/crisszkutnik/discord-notifier/internal/discord"
)

type NotificationService struct {
	discordService *discord.DiscordService
}

func NewNotificationService(discordService *discord.DiscordService) *NotificationService {
	return &NotificationService{
		discordService: discordService,
	}
}

func (s *NotificationService) SendNotification(channelId string, payload *discordgo.MessageSend) error {
	err := s.discordService.SendComplexMessage(channelId, payload)
	return err
}
