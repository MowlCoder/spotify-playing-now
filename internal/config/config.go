package config

import "flag"

type AppConfig struct {
	HttpPort       int
	UpdateInterval int
}

func (cfg *AppConfig) ParseFlags() {
	flag.IntVar(&cfg.HttpPort, "p", 7766, "port number which http handler will be listening to")
	flag.IntVar(&cfg.UpdateInterval, "i", 5, "update current track interval in seconds")

	flag.Parse()
}
