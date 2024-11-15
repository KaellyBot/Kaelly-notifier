package mappers

import (
	"github.com/bwmarrin/discordgo"
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	i18n "github.com/kaysoro/discordgo-i18n"
)

func MapSetNews(setNews *amqp.NewsSetMessage, game amqp.Game,
	locale amqp.Language) *discordgo.WebhookParams {
	lg := constants.MapAMQPLocale(locale)
	return &discordgo.WebhookParams{
		Username:  constants.ExternalName,
		AvatarURL: constants.AvatarURL,
		Content: i18n.Get(lg, "set.message", i18n.Vars{
			"game":             constants.GetGame(game).Name,
			"missingSetNumber": setNews.MissingSetNumber,
			"builtSetNumber":   setNews.BuiltSetNumber,
		}),
	}
}
