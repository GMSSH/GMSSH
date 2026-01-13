package main

import (
	"path/filepath"

	"github.com/DemonZack/simplejrpc-go/core"
	"github.com/DemonZack/simplejrpc-go/core/config"
	"github.com/DemonZack/simplejrpc-go/example/app/server"
	"github.com/DemonZack/simplejrpc-go/os/gpath"
)

func main() {
	env := "test"
	fullPath, _ := filepath.Abs(filepath.Dir("."))
	gpath.GmCfgPath = filepath.Join(filepath.Dir(fullPath), "..")
	core.InitContainer(config.WithConfigEnvFormatterOptionFunc(env))

	appServer := server.NewAppServer()
	appServer.Run()
}
