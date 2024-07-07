package config

import "github.com/spf13/viper"

type Config struct {
	Port string `mapstructure:"PORT"`
}

func LoadConfig() (*Config, error) {
	c := &Config{}

	viper.AddConfigPath("pkg/config")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(c); err != nil {
		return nil, err
	}

	return c, nil
}
