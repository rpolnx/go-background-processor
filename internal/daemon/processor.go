package daemon

import (
	"fmt"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"github.com/rpolnx/go-background-processor/internal/configs"
	"github.com/sirupsen/logrus"
)

type Processor struct {
	workerPool *work.WorkerPool
}

func (p *Processor) ProcessJobs() {
	p.workerPool.Start()
}

func (p *Processor) Stop() {
	p.workerPool.Stop()
}

type ProcessorContext struct {
	customerID int64
}

func (c *ProcessorContext) Log(job *work.Job, next work.NextMiddlewareFunc) error {
	fmt.Println("Starting job: ", job.Name)
	return next()
}

func (c *ProcessorContext) FindCustomer(job *work.Job, next work.NextMiddlewareFunc) error {
	// If there's a customer_id param, set it in the Processorcontext for future middleware and handlers to use.
	if _, ok := job.Args["customer_id"]; ok {
		c.customerID = job.ArgInt64("customer_id")
		if err := job.ArgError(); err != nil {
			return err
		}
	}

	return next()
}

func (c *ProcessorContext) SendEmail(job *work.Job) error {
	logrus.Info("Processing send_welcome_email")
	// Extract arguments:
	addr := job.ArgString("address")
	subject := job.ArgString("subject")
	jobId := job.ArgInt64("job_id")
	if err := job.ArgError(); err != nil {
		return err
	}

	// Go ahead and send the email...
	// sendEmailTo(addr, subject)
	logrus.Warnf("Finishing send_welcome_email -- values to job_id: %d, %s - %s", jobId, addr, subject)

	return nil
}

func (c *ProcessorContext) TriggerJobAtTime(job *work.Job) error {
	logrus.Info("Processing trigger_job_at_time")
	// Extract arguments:
	jobId := job.ArgInt64("job_id")
	userId := job.ArgInt64("user_id")
	timestamp := job.ArgString("timestamp")
	if err := job.ArgError(); err != nil {
		return err
	}
	logrus.Warnf("Finishing trigger_job_at_time -- values to job_id: %d, %d - timestamp: %s", jobId, userId, timestamp)

	return nil
}

func (c *ProcessorContext) Export(job *work.Job) error {
	return nil
}

func NewProcessor(appConfig *configs.AppConfig, cachePool *redis.Pool) *Processor {
	pool := work.NewWorkerPool(ProcessorContext{}, 10, appConfig.AppName, cachePool)

	// Add middleware that will be executed for each job
	pool.Middleware((*ProcessorContext).Log)
	pool.Middleware((*ProcessorContext).FindCustomer)

	pool.Job("send_welcome_email", (*ProcessorContext).SendEmail)
	pool.Job("trigger_job_at_time", (*ProcessorContext).TriggerJobAtTime)

	return &Processor{
		workerPool: pool,
	}
}
