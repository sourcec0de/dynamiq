package app

import "code.google.com/p/gcfg"
import "log"

type Config struct {
	Core struct {
		Name       string
		Port       int
		SeedServer string
		SeedPort   int
		HttpPort   int
		RingSize   int
		Visibility float64
	}
}

func Getconfig(config_file *string) (Config, error) {
	var cfg Config
	err := gcfg.ReadFileInto(&cfg, *config_file)
	if err != nil {
		log.Fatal(err)
	}
	return cfg, err
}