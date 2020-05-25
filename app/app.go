package app

import (
	"fmt"
	"github-integration/app/constants"
	"github-integration/app/model"
	"github-integration/config"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	"log"
	"net/http"
	"os"
	"path/filepath"
	"github.com/spf13/viper"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}
func (a *App) DbInitialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}


	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	set:=&App{a.Router,a.DB}
	a.setRouters(set)
}

func  InitConfig() (err error) {
	path := os.Getenv(constants.ConfigPath)
	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	log.Println("Searching for application configuration file...")
	if path == "" {
		p, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		viper.AddConfigPath(p + constants.ConfigPath)
		log.Println("Loading configs from default location...")
	} else {
		viper.AddConfigPath(path)
		log.Printf("Loading configs from location: %s\n", path)
	}
	err =viper.ReadInConfig()
	if err != nil {
		return err
	}
	return
}

func (a *App) Close() error {
	return a.DB.Close()
}
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}
func (a *App) Run(host string) {
	log.Printf("Starting server at port %v", host)
	log.Fatal(http.ListenAndServe(host, a.Router))
}
