package daemon

import (
	"log"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"github.com/rpolnx/go-background-processor/internal/configs"
)

type Enqueuer struct {
	Enqueuer *work.Enqueuer
}

func (d *Enqueuer) EnqueueJob1() {
	_, err := d.Enqueuer.Enqueue("send_welcome_email", work.Q{"email_address": "test@example.com", "user_id": 4, "customer_id": 1})
	if err != nil {
		log.Fatal(err)
	}
}

func (d *Enqueuer) EnqueueJob2() {
	_, err := d.Enqueuer.Enqueue("trigger_job_at_time", work.Q{"job_id": 1, "user_id": 4, "timestamp": "2024-02-21T02:54:29.587Z", "customer_id": 2})
	if err != nil {
		log.Fatal(err)
	}
}

func NewEnqueuer(appConfig *configs.AppConfig, cachePool *redis.Pool) *Enqueuer {
	var enqueuer = work.NewEnqueuer(appConfig.Host, cachePool)

	return &Enqueuer{
		Enqueuer: enqueuer,
	}
}
