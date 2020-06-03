package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit"
)

type watchBot struct {
	bot reddit.Bot
}

func newReminderBot(bot reddit.Bot) watchBot {
	return watchBot{
		bot: bot,
	}
}

func (r *watchBot) Post(p *reddit.Post) error {
	log.Infof("Post received: %v", p)
	return nil
}


func listen(cfg reddit.BotConfig, subreddits []string) (stop func(), wait func() error, err error) {
	bot, err := reddit.NewBot(cfg)
	if err != nil {
		return nil, nil, err
	}

	settings := graw.Config{
		Subreddits: subreddits,
	}

	handler := newReminderBot(bot)

	return graw.Run(&handler, bot, settings)
}
