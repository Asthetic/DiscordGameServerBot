package discord

import (
	"fmt"

	"github.com/Asthetic/DiscordGameServerBot/config"
	"github.com/bwmarrin/discordgo"
)

type Discord struct {
	session  *discordgo.Session
	token    string
	channels []string
}

// Session is our global discord session
var Session, _ = discordgo.New()

// New creates a new Discord session and saves config
func New(cfg config.Discord) (*Discord, error) {
	dg, err := discordgo.New(fmt.Sprintf("Bot %s", cfg.Token))
	if err != nil {
		return nil, err
	}

	if err = dg.Open(); err != nil {
		return nil, err
	}

	return &Discord{
		session:  dg,
		token:    cfg.Token,
		channels: cfg.Channels,
	}, nil
}

// Close cleanly closes the current Discord session
func (d *Discord) Close() {
	d.session.Close()
}

// SendUpdatedIP sends the updated IP address to the configured channels
func (d *Discord) SendUpdatedIP(ip string) {
	for _, channel := range d.channels {
		msg := formatMsg(ip)
		d.session.ChannelMessageSendComplex(channel, msg)
	}
}

func formatMsg(ip string) *discordgo.MessageSend {
	return &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Type:  discordgo.EmbedTypeRich,
			Title: "Minecraft Server IP Updated",
			Color: 5439264,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL:    "https://i.imgur.com/6rp13.png",
				Width:  128,
				Height: 128,
			},
			Fields: formatFields(ip),
		},
	}
}

func formatFields(ip string) []*discordgo.MessageEmbedField {
	fields := []*discordgo.MessageEmbedField{}
	field := &discordgo.MessageEmbedField{
		Name:   "Minecraft",
		Value:  ip,
		Inline: true,
	}

	fields = append(fields, field)
	return fields
}
