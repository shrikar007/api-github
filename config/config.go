package config

import "github.com/spf13/viper"

type Config struct {
	DB *DBConfig
}
 
type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
}
 
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  viper.GetString("app.dialect"),
			Username: viper.GetString("app.username"),
			Password: viper.GetString("app.password"),
			Name:     viper.GetString("app.name"),
			Charset:  viper.GetString("app.charset"),
		},
	}
}
