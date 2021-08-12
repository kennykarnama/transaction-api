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
	"transaction-api/domain/repository/transaction/mysql"
	transService "transaction-api/domain/service/transaction"
	"transaction-api/util/dbconn"
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

	db := dbconn.InitGorm(cfg.ServiceName)

	transRepo := mysql.NewRepository(db)

	transService := transService.NewService(transRepo)
	transHandler := transaction.NewHandler(ctx, v, transService)
	httpServer := &server{
		Server: http.Server{
			Addr: ":" + cfg.RestPort,
		},
	}
	r := mux.NewRouter()
	r.Handle("/api/v1/order", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(transHandler.CreateTransaction))).Methods("POST")
	r.Handle("/api/v1/order/withPayment", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(transHandler.CreateAndPayTransaction))).Methods("POST")
	r.Handle("/api/v1/orders/user/{id}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(transHandler.ListUserTransaction))).Methods("GET")
	r.Handle("/api/v1/order/{id}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(transHandler.GetTransactionDetail))).Methods("GET")
	r.Handle("/api/v1/order/{id}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(transHandler.UpdateTransactionByID))).Methods("PUT")
	r.Handle("/api/v1/order/{id}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(transHandler.DeleteTransactionByID))).Methods("DELETE")
	r.Handle("/api/v1/order/{id}/items", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(transHandler.DeleteTransactionItems))).Methods("DELETE")

	httpServer.Handler = r

	logrus.WithFields(standardFields).Infof("HTTP served on port: %v", cfg.RestPort)

	if err := httpServer.ListenAndServe(); err != nil {
		logrus.WithFields(standardFields).Fatalf("unable to serve. err: %v", err)
	}
}
