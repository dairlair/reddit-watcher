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
	postsRepository postsRepositoryInterface
}

func newWatcherComponent(
		cfg reddit.BotConfig,
		subreddits []string,
		postsRepository postsRepositoryInterface,
	) watcherComponent {
	return watcherComponent{
		cfg,
		subreddits,
		postsRepository,
	}
}

func (w watcherComponent) Run() (stop func(), wait func() error, err error) {
	return listen(w.cfg, w.subreddits, w.postsRepository)
}

func (w watcherComponent) IsReady() bool {
	return true
}

func main() {
	repo := newMongoRepository("mongodb://localhost:27017")

	subreddits := strings.Split(os.Getenv("REDDIT_SUBREDDITS"), " ")
	watcher := newWatcherComponent(getBotConfig(), subreddits, repo)

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