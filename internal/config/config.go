package config

import (
	"github.com/spf13/viper"
)

type NSQConfig struct {
	Host     string `mapstructure:"HOST"`
	Port     string `mapstructure:"PORT"`
	Topic    string `mapstructure:"TOPIC"`
	Channel  string `mapstructure:"CHANNEL"`
	Topic2   string `mapstructure:"TOPIC2"`
	Channel2 string `mapstructure:"CHANNEL2"`
}

type SenderConfig struct {
	Email    string `mapstructure:"EMAIL"`
	Password string `mapstructure:"PASSWORD"`
}

type Config struct {
	Sender SenderConfig `mapstructure:"SENDER"`
	NSQ    NSQConfig    `mapstructure:"NSQ"`
}

func NewConfig() (*Config, error) {
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	viper.AutomaticEnv()

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
