package config

import "github.com/spf13/viper"

type Config struct {
	Port       int `mapstructure:"port"`
	PostgreSQL `mapstructure:"postgres"`
}

type PostgreSQL struct {
	URI      string `mapstructure:"uri"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database_name"`
}

func LoadConfig() (config Config, err error) {

	//Initialize properties config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return config, err
}
