package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/rpolnx/go-background-processor/internal/configs"
	"github.com/rpolnx/go-background-processor/internal/daemon"
	handler "github.com/rpolnx/go-background-processor/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	appConfig, err := configs.InitEnvVariables()
	if err != nil {
		logrus.Fatalln(err)
	}

	configs.InitLogger()

	cachePool := configs.NewCachePool()

	server, err := handler.InitializeServer(appConfig)

	if err != nil {
		log.Fatal("Error initializing server", err)
	}

	enqueuer := daemon.NewEnqueuer(appConfig, cachePool)
	//continue create jobs

	go func() {
		rand.Seed(time.Now().UnixNano())

		for idx := 0; ; idx++ {
			time.Sleep(time.Second * time.Duration(5))
			enqueuer.EnqueueJob1(idx)
			enqueuer.EnqueueJob2(idx)

			randomInt := rand.Intn(101)
			enqueuer.RegisterScheduledJob(idx, randomInt)

			logrus.Info("hey ooo")
		}
	}()

	processor := daemon.NewProcessor(appConfig, cachePool)
	go processor.ProcessJobs()

	// signalChan := make(chan os.Signal, 1)
	// signal.Notify(signalChan, os.Interrupt, os.Kill)
	// <-signalChan

	serverHost := fmt.Sprintf("%s:%d", configs.GlobalAppConfig.Host, configs.GlobalAppConfig.Port)
	if err = server.Engine.Run(serverHost); err != nil {
		log.Fatalln(err)
	}

	processor.Stop()
}
