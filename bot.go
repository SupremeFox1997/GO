package bot

import (
	"github.com/bwmarrin/discordgo"
)

// Глобальная переменная для сессии бота
var botSession *discordgo.Session

type BotSession struct {
	*discordgo.Session
}

func NewSession(token string) (*BotSession, error) {
	session, err := discordgo.New(token)
	if err != nil {
		return nil, err
	}
	return &BotSession{Session: session}, nil
}

func (b *BotSession) Open() error {
	return b.Session.Open()
}

func (b *BotSession) Close() {
	b.Session.Close()
}

func (b *BotSession) CreateChannel(name, categoryID string) (*discordgo.Channel, error) {
	return b.Session.GuildChannelCreateComplex(ServerID, discordgo.GuildChannelCreateData{
		Name:     name,
		Type:     discordgo.ChannelTypeGuildVoice,
		ParentID: categoryID,
	})
}

func (b *BotSession) MoveUserToChannel(userID, channelID string) error {
	return b.Session.GuildMemberMove(ServerID, userID, &channelID)
}

func (b *BotSession) DeleteChannel(channelID string) error {
	return b.Session.ChannelDelete(channelID)
}
