package config

import "flag"

type Config struct {
	Version   string
	Env       string
	StaticDir string
	Port      int
}

var version = "0.0.1"

func LoadConfig() Config {
	var cfg Config

	flag.StringVar(&cfg.Version, "version", version, "Application version")
	flag.StringVar(&cfg.Env, "env", "development", "Enviroment (development, staging, production)")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.IntVar(&cfg.Port, "port", 8080, "Webserver port")

	flag.Parse()

	return cfg
}
