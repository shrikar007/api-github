package app

import (
	"fmt"
	"github-integration/app/model"
	"github-integration/drivers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}
func (a *App) DbInitialize(config *drivers.Config) {
     var dbURI string
	if config.DB.Dialect=="mysql"{
		dbURI = fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
			config.DB.Username,
			config.DB.Password,
			config.DB.Name,
			config.DB.Charset)
	}else{
		dbURI=fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
			config.DB.Host,
			config.DB.Port,
			config.DB.Username,
			config.DB.Name,
			config.DB.Password)
	}
	db , err :=gorm.Open(config.DB.Dialect,dbURI)
	if err != nil {
		log.Fatal("Could not connect database",err)
	}
	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	set:=&App{a.Router,a.DB}
	a.setRouters(set)
}
func Initlogger(){
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logrus.SetLevel(logrus.DebugLevel)
}

func  InitConfig() (err error) {
	viper.SetConfigType("toml")
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")
	err =viper.ReadInConfig()
	if err != nil {
		logrus.WithError(err).Error("unable read config file")
	}
	return
}

func (a *App) Close() error {
	return a.DB.Close()
}
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")

}
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}
func (a *App) Run(host string) {
	logrus.Printf("Starting server at port %v", host)
	logrus.Fatal(http.ListenAndServe(host, a.Router))
}
