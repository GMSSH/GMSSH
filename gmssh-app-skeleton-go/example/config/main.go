package main

import (
	"fmt"
	"path/filepath"

	"github.com/DemonZack/simplejrpc-go/core"
	"github.com/DemonZack/simplejrpc-go/core/config"
	"github.com/DemonZack/simplejrpc-go/os/gpath"
)

func main() {
	env := "test"
	fullPath, _ := filepath.Abs(filepath.Dir("."))
	gpath.GmCfgPath = filepath.Join(filepath.Dir(fullPath), "..")
	core.InitContainer(config.WithConfigEnvFormatterOptionFunc(env))

	val, err := core.Container.CfgFmt().GetValue("logger.level").String()
	if err != nil {
		panic(err)
	}
	fmt.Println("[*] logger.level : ", val)

	val, err = core.Container.CfgFmt().GetValue("jsonrpc.sockets").String()
	if err != nil {
		panic(err)
	}
	fmt.Println("[*] jsonrpc.sockets : ", val)
}
