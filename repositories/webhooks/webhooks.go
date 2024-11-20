package webhooks

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-notifier/models/constants"
	"github.com/kaellybot/kaelly-notifier/models/entities"
	"github.com/kaellybot/kaelly-notifier/utils/databases"
	"gorm.io/gorm"
)

func New(db databases.MySQLConnection) *Impl {
	return &Impl{db: db}
}

func (repo *Impl) GetAlmanaxWebhooks(game amqp.Game, locale amqp.Language,
) ([]*entities.WebhookAlmanax, error) {
	var webhooks []*entities.WebhookAlmanax
	err := repo.db.GetDB().
		Where("game = ? AND locale = ?", game, locale).
		Find(&webhooks).Error
	if err != nil {
		return nil, err
	}

	return webhooks, nil
}

func (repo *Impl) GetFeedWebhooks(feedTypeID string, game amqp.Game, locale amqp.Language,
) ([]*entities.WebhookFeed, error) {
	var webhooks []*entities.WebhookFeed
	err := repo.db.GetDB().
		Where("feed_type_id = ? AND game = ? AND locale = ?",
			feedTypeID, game, locale).
		Find(&webhooks).Error
	if err != nil {
		return nil, err
	}

	return webhooks, nil
}

func (repo *Impl) GetTwitchWebhooks(streamerID string) ([]*entities.WebhookTwitch, error) {
	var webhooks []*entities.WebhookTwitch
	err := repo.db.GetDB().
		Where("streamer_id = ?", streamerID).
		Find(&webhooks).Error
	if err != nil {
		return nil, err
	}

	return webhooks, nil
}

func (repo *Impl) GetTwitterWebhooks(twitterID string) ([]*entities.WebhookTwitter, error) {
	var webhooks []*entities.WebhookTwitter
	err := repo.db.GetDB().
		Where("twitter_id = ?", twitterID).
		Find(&webhooks).Error
	if err != nil {
		return nil, err
	}

	return webhooks, nil
}

func (repo *Impl) GetYoutubeWebhooks(videastID string) ([]*entities.WebhookYoutube, error) {
	var webhooks []*entities.WebhookYoutube
	err := repo.db.GetDB().
		Where("videast_id = ?", videastID).
		Find(&webhooks).Error
	if err != nil {
		return nil, err
	}

	return webhooks, nil
}

func (repo *Impl) UpdateWebhooks(model any, webhooks []*constants.Webhook) error {
	return repo.db.GetDB().Transaction(func(tx *gorm.DB) error {
		for _, wh := range webhooks {
			err := tx.Model(model).
				Where("webhook_id = ?", wh.WebhookID).
				Updates(map[string]any{
					"retry_number": wh.RetryNumber,
					"failed_at":    wh.FailedAt,
					"published_at": wh.PublishedAt,
				}).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (repo *Impl) DeleteWebhooks(model any, webhooks []*constants.Webhook) error {
	return repo.db.GetDB().Transaction(func(tx *gorm.DB) error {
		for _, wh := range webhooks {
			err := tx.Where("webhook_id = ?", wh.WebhookID).Delete(model).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}