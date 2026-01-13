package main

import (
	"fmt"
	"path/filepath"

	"github.com/DemonZack/simplejrpc-go/core"
	"github.com/DemonZack/simplejrpc-go/core/config"
	"github.com/DemonZack/simplejrpc-go/os/gpath"
)

type ExampleUser struct {
	Username string `validate:"min_length:6#The length is too small"`
	Age      any    `validate:"required#Required parameters are missing Age|range:18,100|int#Test verification error return"`
	Email    string `validate:"required#Email address is required"`
}

func main() {
	env := "test"
	fullPath, _ := filepath.Abs(filepath.Dir("."))
	gpath.GmCfgPath = filepath.Join(filepath.Dir(fullPath), "..")
	core.InitContainer(config.WithConfigEnvFormatterOptionFunc(env))

	user := ExampleUser{
		Username: "test",
		Age:      15,
		Email:    "",
	}

	err := core.Container.Valid().Walk(&user)
	if err != nil {
		fmt.Println("Validation error:", err)
	}
}
