package main

import (
	"github.com/andreybutko/payment/usecase/payment"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/andreybutko/payment/api/handler"
	"github.com/andreybutko/payment/config"
)

func main() {
	r := mux.NewRouter()

	paymentService := payment.NewService()
	handler.MakePaymentHandlers(r, paymentService)

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(config.API_PORT),
		Handler:      http.DefaultServeMux,
		ErrorLog:     logger,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}