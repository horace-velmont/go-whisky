package main

import (
	"github.com/GagulProject/go-whisky/server"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		server.InvokeServer(),
		server.Options(),
	).Run()
}
