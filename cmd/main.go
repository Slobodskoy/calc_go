package main

import (
	"log"

	"github.com/Slobodskoy/calc_go/internal/app"
	"github.com/Slobodskoy/calc_go/internal/config"
	"github.com/caarlos0/env/v11"
)

func main() {
	var cfg config.Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	app := app.New(cfg)
	app.Run()
}
