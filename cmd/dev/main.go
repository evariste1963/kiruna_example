package main

import (
	"kiruna_example/internal/platform"

	"github.com/sjc5/kiruna"
)

func main() {
	platform.Kiruna.MustStartDev(&kiruna.DevConfig{
		HealthcheckEndpoint: "/healthz",
		WatchedFiles:        kiruna.WatchedFiles{{Pattern: "**/*.go.html"}},
	})
}
