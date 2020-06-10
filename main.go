package main

import (
	"github.com/dairlair/reddit-watcher/pkg/storage"
	"github.com/dairlair/reddit-watcher/pkg/watcher"
	"github.com/stepsisters/kgb"
	"github.com/stepsisters/kgb/pkg/component/kubernetes"
	"github.com/stepsisters/kgb/pkg/component/signal"
	"os"
)

func main() {
	s := storage.NewStorageComponent()
	w := watcher.NewWatcherComponent(&s)

	probesPort := os.Getenv("PROBES_PORT")
	probe := kubernetes.NewHTTPProbe(w.IsReady, probesPort)

	components := map[string]kgb.ComponentInterface{
		"k8s-probes": probe,
		"signals-trap": signal.NewTrap(),
		"storage":  &s,
		"watcher":  &w,
	}
	kgb.Run(components)
}