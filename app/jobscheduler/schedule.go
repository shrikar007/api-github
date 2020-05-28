package jobscheduler

import (
	"context"
	"github.com/zhashkevych/scheduler"
	"time"
)

func Curlresquest(ctx context.Context) {

}
func Jobschedule(){
	ctx := context.Background()
	worker := scheduler.NewScheduler()
	worker.Add(ctx, Curlresquest, time.Second*5)
}
