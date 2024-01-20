package main

import (
	"github.com/prophet0fregret/go-microservice/internal/database"
	"github.com/prophet0fregret/go-microservice/internal/server"
	"github.com/sirupsen/logrus"
)

func init() {
	err := database.InitDatabaseClient()
	if err != nil {
		logrus.WithError(err).Fatal("Unable to initialize database Client.....")
		return
	}
}

func main() {
	//Fetch initialized db client
	dbClient, err := database.ReturnDatabaseClient()
	if err != nil {
		logrus.WithError(err).Fatal("Unable to initialize database Client.....")
		return
	}

	//Create new echo server
	srv := server.NewEchoServer(dbClient)
	if err = srv.Start(); err != nil {
		logrus.WithError(err).Fatal("Unable to start Echo server.....")
		return
	}

	logrus.Info("Started server & Listening to incoming requests......")
}
