package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DB *DBConfig
}
 
type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
	Port     string
	Host     string
}
 
func GetConfig() *Config {
	if viper.GetBool("app.DB"){
		return &Config{
			DB: &DBConfig{
				Dialect:  viper.GetString("database.mysql.dialect"),
				Username: viper.GetString("database.mysql.username"),
				Password: viper.GetString("database.mysql.password"),
				Name:     viper.GetString("database.mysql.name"),
				Charset:  viper.GetString("database.mysql.charset"),
			},
		}

	} else{
		return &Config{
			DB: &DBConfig{
				Dialect:  viper.GetString("database.postgres.dialect"),
				Username: viper.GetString("database.postgres.username"),
				Password: viper.GetString("database.postgres.password"),
				Name:     viper.GetString("database.postgres.name"),
				Port: viper.GetString("database.postgres.port"),
				Host: viper.GetString("database.postgres.host"),
			},
		}
	}
	return nil

}
