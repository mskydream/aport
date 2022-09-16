package config

import "github.com/spf13/viper"

type Config struct {
	DatabaseSource string `mapstructure:"dbSource"`
	Port           string `mapstructure:"port"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
