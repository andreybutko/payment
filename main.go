package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/andreybutko/payment/api/handler"
	"github.com/andreybutko/payment/api/middleware"
	"github.com/andreybutko/payment/configuration"
	"github.com/andreybutko/payment/pkg/metric"
	"github.com/andreybutko/payment/usecase/payment"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	provider := payment.NewProvider()
	paymentService := payment.NewService(provider)
	metricService, err := metric.NewMetricService()
	n := negroni.New(
		middleware.Metrics(metricService),
		middleware.Common(),
	)
	handler.MakePaymentHandlers(r, *n, paymentService)
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	http.Handle("/", r)

	config := configuration.GetConfig()
	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + config.HostPort,
		Handler:      http.DefaultServeMux,
		ErrorLog:     logger,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
