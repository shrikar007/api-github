package jobscheduler

import (
	"github.com/spf13/viper"
	"gopkg.in/robfig/cron.v2"
	"net/http"
)

func Curlresquest1() {
	url := viper.GetString("app.schedule_url")
	req, _ := http.NewRequest("GET", url, nil)
	_, _ = http.DefaultClient.Do(req)
}
func Curlresquest2() {
	url := viper.GetString("app.schedule_url")
	req, _ := http.NewRequest("GET", url, nil)
	_, _ = http.DefaultClient.Do(req)
}
func Curlresquest3() {
	url := viper.GetString("app.schedule_url")
	req, _ := http.NewRequest("GET", url, nil)
	_, _ = http.DefaultClient.Do(req)
}
func Noon(){
	c := cron.New()
	c.AddFunc("TZ=Asia/Kolkata 14 54 * * * *", Curlresquest1)
	c.AddFunc("@hourly", Curlresquest2)
	c.AddFunc("@every 0h0m1s", Curlresquest3)
	c.Start()
}
