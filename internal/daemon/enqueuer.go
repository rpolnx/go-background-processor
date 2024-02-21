package daemon

import (
	"log"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"github.com/rpolnx/go-background-processor/internal/configs"
	"github.com/sirupsen/logrus"
)

type Enqueuer struct {
	Enqueuer *work.Enqueuer
}

func (d *Enqueuer) EnqueueJob1(idx int) {
	logrus.Infof("Enqueued send_welcome_email - %d", idx)
	_, err := d.Enqueuer.Enqueue("send_welcome_email",
		work.Q{
			"job_id":        idx,
			"email_address": "test@example.com",
			"user_id":       4,
			"customer_id":   1,
		})
	if err != nil {
		log.Fatal(err)
	}
}

func (d *Enqueuer) EnqueueJob2(idx int) {
	logrus.Infof("Enqueued trigger_job_at_time - %d", idx)
	_, err := d.Enqueuer.Enqueue("trigger_job_at_time",
		work.Q{
			"job_id":      idx,
			"user_id":     4,
			"timestamp":   "2024-02-21T02:54:29.587Z",
			"customer_id": 2,
		})
	if err != nil {
		log.Fatal(err)
	}
}

func NewEnqueuer(appConfig *configs.AppConfig, cachePool *redis.Pool) *Enqueuer {
	var enqueuer = work.NewEnqueuer(appConfig.AppName, cachePool)

	return &Enqueuer{
		Enqueuer: enqueuer,
	}
}
