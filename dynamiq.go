package main

import (
	"flag"

	"github.com/Sirupsen/logrus"
	"github.com/Tapjoy/dynamiq/app"
)

func main() {
	//Get some Command line options
	configFile := flag.String("c", "./lib/config.gcfg", "location of config file")
	flag.Parse()

	if *configFile == "" {
		logrus.Warn("Empty value provided for config file location from flag -c : Falling back to default location './lib/config.gcfg'")
		*configFile = "./lib/config.gcfg"
	}

	//setup the config file
	cfg, err := app.GetCoreConfig(configFile)

	cfg.Topics = app.InitTopics(cfg, cfg.Queues)

	if err != nil {
		logrus.Fatal(err)
	}
	logrus.SetLevel(cfg.Core.LogLevel)

	list, _, err := app.InitMemberList(cfg.Core.Name, cfg.Core.Port, cfg.Core.SeedServers, cfg.Core.SeedPort)
	httpAPI := app.HTTPApiV1{}

	httpAPI.InitWebserver(list, cfg)
}
