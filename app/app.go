package app

import (
	"fmt"
	"github-integration/app/jobscheduler"
	"github-integration/app/model"
	"github-integration/drivers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}
func New()(apps App,err error){
	app:=App{}
	Initlogger()
	logrus.Print("Logger initialized")
	logrus.Print("Loading configuration")
	if err=InitConfig();err!=nil{
		return
	}
	logrus.Print("Configuration Loaded")
	configObj := drivers.GetConfig()
	logrus.Print("Start Application")
	db:=app.DbInitialize(configObj)
	route:=app.Initroutes()
	set:=&App{route,db}
	app.setRouters(set)
	return app, nil
}
func (a *App) DbInitialize(config *drivers.Config) (*gorm.DB) {
     var dbURI string
	if config.DB.Dialect=="mysql"{
		dbURI = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True",
			config.DB.Username,
			config.DB.Password,
			config.DB.Host,
			config.DB.Port,
       		config.DB.Name,
			config.DB.Charset,)
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
		logrus.WithError(err).Error("Could not connect database")
	}
	a.DB = model.DBMigrate(db)
	return a.DB
}
func(a *App)  Initroutes() (*mux.Router) {
	a.Router=mux.NewRouter()
	return a.Router
}
func Initlogger(){
	homeDirPath, err := os.UserHomeDir()
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	runID := time.Now().Format("git-api-2006-01-02-15-04-05")
	_, err = os.Stat(path.Join(homeDirPath, "git-api-app","logs"))
	if err != nil {
		err = os.MkdirAll(path.Join(homeDirPath, "git-api-app","logs"), os.ModePerm)
		if err != nil {
			logrus.WithError(err).Error("unable to create logs folder for app")
		}
	}
	logLocation := filepath.Join(homeDirPath,"git-api-app","logs", runID + ".log")
	logFile, err := os.OpenFile(logLocation, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logrus.WithError(err).Error("Failed to open log file")
	}
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(io.MultiWriter(os.Stderr, logFile))
}
func  InitConfig() (err error) {
	viper.SetConfigType("toml")
	viper.SetConfigName("config")
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
func (a *App) Run(host string) {
	logrus.Printf("Starting server at port %v", host)
	go jobscheduler.Jobschedule()
	logrus.Fatal(http.ListenAndServe(host, a.Router))
}
