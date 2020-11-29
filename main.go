package main

import (
	"github.com/andreybutko/payment/api/handler"
	"github.com/andreybutko/payment/api/middleware"
	"github.com/andreybutko/payment/config"
	"github.com/andreybutko/payment/pkg/metric"
	"github.com/andreybutko/payment/usecase/payment"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	r := mux.NewRouter()

	paymentService := payment.NewService()
	metricService, err := metric.NewMetricService()
	n := negroni.New(
		middleware.Metrics(metricService),
	)
	handler.MakePaymentHandlers(r, *n, paymentService)
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	http.Handle("/", r)

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(config.API_PORT),
		Handler:      http.DefaultServeMux,
		ErrorLog:     logger,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}