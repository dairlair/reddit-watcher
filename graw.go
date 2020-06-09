package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit"
)

type watchBot struct {
	bot reddit.Bot
	postsRepository postsRepositoryInterface
}

func newReminderBot(bot reddit.Bot, postsRepository postsRepositoryInterface) watchBot {
	return watchBot{
		bot,
		postsRepository,
	}
}

func (r *watchBot) Post(p *reddit.Post) error {
	log.Infof("Post received: %v", p)
	return r.postsRepository.save(*p)
}

type postsRepositoryInterface interface {
	save (post reddit.Post) error
}

func listen(
		cfg reddit.BotConfig,
		subreddits []string,
		postsRepository postsRepositoryInterface,
	) (stop func(), wait func() error, err error) {
	bot, err := reddit.NewBot(cfg)
	if err != nil {
		return nil, nil, err
	}

	settings := graw.Config{
		Subreddits: subreddits,
	}

	handler := newReminderBot(bot, postsRepository)

	return graw.Run(&handler, bot, settings)
}
