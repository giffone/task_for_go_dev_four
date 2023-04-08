package main

import (
	"bot_svc/internal/app"
	"bot_svc/internal/config"
)

func main() {
	// configuration
	conf := config.NewConf()

	// application
	a := app.NewApp(conf)
	a.ListenAndServe(conf)
	
}
