package main

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"transaction-api/config"
	"transaction-api/domain/api/transaction"
)

type server struct {
	http.Server
}

func main() {
	ctx := context.Background()
	cfg := config.Get()

	hostName, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	logrus.SetFormatter(&logrus.JSONFormatter{})
	standardFields := logrus.Fields{
		"appname":  cfg.ServiceName,
		"hostname": hostName,
	}
	v := validator.New()

	transHandler := transaction.NewHandler(ctx, v, nil)
	httpServer := &server{
		Server: http.Server{
			Addr: ":" + cfg.RestPort,
		},
	}
	r := mux.NewRouter()
	r.Handle("/api/v1/order", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(transHandler.CreateTransaction))).Methods("POST")

	httpServer.Handler = r

	logrus.WithFields(standardFields).Infof("HTTP served on port: %v", cfg.RestPort)

	if err := httpServer.ListenAndServe(); err != nil {
		logrus.WithFields(standardFields).Fatalf("unable to serve. err: %v", err)
	}
}
