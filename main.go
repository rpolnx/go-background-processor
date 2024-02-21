package main

import (
	"fmt"
	"log"

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
	enqueuer.EnqueueJob1()
	enqueuer.EnqueueJob2()

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
