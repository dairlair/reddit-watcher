package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"github.com/stepsisters/kgb"
	"github.com/stepsisters/kgb/pkg/component/kubernetes"
	"github.com/stepsisters/kgb/pkg/component/signal"
	"github.com/turnage/graw/reddit"
	"os"
	"strings"
)

type watcherComponent struct {
	cfg reddit.BotConfig
	subreddits []string
}

func newWatcherComponent(cfg reddit.BotConfig, subreddits []string) watcherComponent {
	return watcherComponent{
		cfg,
		subreddits,
	}
}

func (w watcherComponent) Run() (stop func(), wait func() error, err error) {
	return listen(w.cfg, w.subreddits)
}

func (w watcherComponent) IsReady() bool {
	return true
}

func main() {
	subreddits := strings.Split(os.Getenv("REDDIT_SUBREDDITS"), " ")
	watcher := newWatcherComponent(getBotConfig(), subreddits)

	probesPort := os.Getenv("PROBES_PORT")
	probe := kubernetes.NewHTTPProbe(watcher.IsReady, probesPort)

	components := map[string]kgb.ComponentInterface{
		"k8s-probes": probe,
		"signals-trap": signal.NewTrap(),
		"reddit-watcher":  watcher,
	}
	kgb.Run(components)
}

func getBotConfig() reddit.BotConfig {
	flag.Parse()
	cfg := reddit.BotConfig{
		Agent: os.Getenv("REDDIT_USER_AGENT"),
		App: reddit.App{
			ID:       os.Getenv("REDDIT_CLIENT_ID"),
			Secret:   os.Getenv("REDDIT_SECRET"),
			Username: os.Getenv("REDDIT_USERNAME"),
			Password: os.Getenv("REDDIT_PASSWORD"),
		},
		Rate: 0,
	}
	log.Infof("Config: %v", cfg)
	return cfg
}