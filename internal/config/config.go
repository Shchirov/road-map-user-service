package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"sync"
)

type Config struct {
	Env       string `yaml:"env" env-required:"true"`
	AppConfig App    `yaml:"app" env-required:"true"`
}

type App struct {
	HttpConfig     HttpConfig     `yaml:"http" env-required:"true"`
	GrpcConfig     GrpcConfig     `yaml:"grpc" env-required:"true"`
	DataBaseConfig DataBaseConfig `yaml:"db" env-required:"true"`
}

type HttpConfig struct {
	HttpHost string `yaml:"host" env-default:"localhost"`
	HttpPort int    `yaml:"port" env-default:"8080"`
}
type GrpcConfig struct {
	GrpcHost string `yaml:"host" env-default:"localhost"`
	GrpcPort int    `yaml:"port" env-default:"3333"`
}
type DataBaseConfig struct {
	PgUser       string `yaml:"username" env-default:"root"`
	PgPassword   string `yaml:"password" env-default:"root"`
	PostgresDb   string `yaml:"database" env-default:"road-map"`
	PostgresHost string `yaml:"host" env-default:"localhost"`
	PostgresPort int    `yaml:"port" env-default:"5432"`
}

var (
	config *Config
	once   sync.Once
)

// MustLoad Get reads config from environment. Once.
func MustLoad() *Config {
	once.Do(func() {
		configPath := fetchConfigPath()
		if configPath == "" {
			configPath = "configs/config-local.yml"
		}

		config = MustLoadPath(configPath)

	})
	return config
}

func MustLoadPath(configPath string) *Config {
	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}

// fetchConfigPath fetches config path from command line flag or environment variable.
// Priority: flag > env > default.
// Default value is empty string.
func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
