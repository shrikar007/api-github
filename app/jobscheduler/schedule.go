package jobscheduler

import (
	"context"
	"github.com/spf13/viper"
	"github.com/zhashkevych/scheduler"
	"net/http"
	"time"
)

func Curlresquest(ctx context.Context) {
	time.Sleep(time.Second * 1)
	url := viper.GetString("app.schedule_url")
	req, _ := http.NewRequest("GET", url, nil)
	_, _ = http.DefaultClient.Do(req)
}
func Jobschedule(){
	ctx := context.Background()
	worker := scheduler.NewScheduler()
	worker.Add(ctx, Curlresquest, time.Second*5)
}