package config

import (
	"fmt"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type Config struct {
	MaxHubs           uint   `env:"AGORA_MAX_HUBS" envDefault:"1000"`
	MaxChannelsPerHub uint   `env:"AGORA_MAX_CHANNELS_PER_HUB" envDefault:"10"`
	ApiHost           string `env:"AGORA_API_HOST" envDefault:":3000"`
	ApiKey            string `env:"AGORA_API_KEY" envDefault:"secret"`
	DiscordToken      string `env:"DISCORD_TOKEN" envDefault:""`
}

func NewFromEnv(files ...string) (Config, error) {
	config := Config{}

	if errLoad := godotenv.Load(files...); errLoad != nil {
		return Config{}, fmt.Errorf("could not load configuration: %w", errLoad)
	}

	if errParse := env.Parse(&config); errParse != nil {
		return Config{}, fmt.Errorf("could not parse configuration: %w", errParse)
	}

	return config, nil
}
