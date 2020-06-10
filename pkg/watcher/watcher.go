package watcher

import (
	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit"
	"os"
	"strings"
)

type storageInterface interface {
	SavePost (post reddit.Post) error
}

type config struct {
	botConfig      reddit.BotConfig
	watchingConfig graw.Config
}

// WatcherComponent listens reddit and saves received posts to the storage
type WatcherComponent struct {
	config  config
	storage storageInterface
}

func NewWatcherComponent(storage storageInterface) WatcherComponent {
	config := readConfig()
	return WatcherComponent{config: config, storage: storage}
}

// Run function implements the `github.com/stepsisters/kgb.ComponentInterface`
func (w *WatcherComponent) Run () (stop func(), wait func() error, err error) {
	bot, err := reddit.NewBot(w.config.botConfig)
	if err != nil {
		return nil, nil, err
	}

	return graw.Run(w, bot, w.config.watchingConfig)
}

func (w *WatcherComponent) IsReady () bool {
	return true
}

// Post function implements the PostHandler interface
func (w *WatcherComponent) Post (post *reddit.Post) error {
	return w.storage.SavePost(*post)
}

func readConfig() config {
	return config{
		botConfig:  reddit.BotConfig{
			Agent: os.Getenv("REDDIT_USER_AGENT"),
			App: reddit.App{
				ID:       os.Getenv("REDDIT_CLIENT_ID"),
				Secret:   os.Getenv("REDDIT_SECRET"),
				Username: os.Getenv("REDDIT_USERNAME"),
				Password: os.Getenv("REDDIT_PASSWORD"),
			},
			Rate: 0,
		},
		watchingConfig: graw.Config{
			Subreddits: strings.Split(os.Getenv("REDDIT_SUBREDDITS"), " "),
		},
	}
}