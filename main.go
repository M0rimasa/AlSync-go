package main

import (
	"log"

	"github.com/M0rimasa/AlSync-go/config"
	"github.com/M0rimasa/AlSync-go/site"
)



func main() {
	cfgPath, err := config.ParseFlags()
	if err != nil {
			log.Fatal(err)
	}
	Cfg, err := config.NewConfig(cfgPath)
	if err != nil {
			log.Fatal(err)
	}

	site.Updater(*Cfg)
}