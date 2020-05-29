package jobscheduler

import (
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func Curlresquest() {
	url := viper.GetString("app.schedule_url")
	req, _ := http.NewRequest("GET", url, nil)
	_, _ = http.DefaultClient.Do(req)
}

func Jobschedule(){
	done := make(chan bool)
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for {
			select {
			case <-done:
				ticker.Stop()
				return
			case <-ticker.C:
				Curlresquest()
			}
		}
	}()
	time.Sleep(20 *time.Second)
	done <- true
}
func Noon(){
	t := time.Now()
	n := time.Date(t.Year(), t.Month(), t.Day(), 13, 0, 0, 0, t.Location())
	d := n.Sub(t)
	if d < 0 {
		n = n.Add(24 * time.Hour)
		d = n.Sub(t)
	}
	for {
		time.Sleep(d)
		d = 24 * time.Hour
		Jobschedule()
	}

}
