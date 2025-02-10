package platform

import (
	"kiruna_example/dist"

	"github.com/sjc5/kiruna"
)

var Kiruna = kiruna.New(&kiruna.Config{
	DistFS:           dist.FS,
	MainAppEntry:     "cmd/app/main.go",
	PrivateStaticDir: "./static/private",
	PublicStaticDir:  "./static/public",
	StylesDir:        "./styles",
	DistDir:          "./dist",
})
