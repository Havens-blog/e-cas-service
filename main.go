package main

import (
	"github.com/Havens-blog/e-cas-service/cmd"
)

func main() {
	app := cmd.NewApp()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
