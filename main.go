package main

import (
	"fmt"
	"log"

	"github.com/rpolnx/go-background-processor/internal/configs"
	handler "github.com/rpolnx/go-background-processor/internal/server"
	"github.com/sirupsen/logrus"
)

func init() {
	_, err := configs.InitEnvVariables()
	if err != nil {
		logrus.Fatalln(err)
	}

	configs.InitLogger()
}

func main() {
	server, err := handler.InitializeServer(configs.GlobalAppConfig)

	if err != nil {
		log.Fatal("Error initializing server", err)
	}

	err = server.Run(fmt.Sprintf("%s:%d", configs.GlobalAppConfig.Host, configs.GlobalAppConfig.Port))

	if err != nil {
		log.Fatalln(err)
	}
}
