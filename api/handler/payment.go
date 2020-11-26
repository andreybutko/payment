package handler

import (
	"encoding/json"
	"github.com/andreybutko/payment/usecase/payment"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func createPayment(service payment.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding payment"
		var input struct {
			ProductId string `json:"productId"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

//MakeUserHandlers make url handlers
func MakePaymentHandlers(r *mux.Router, service payment.UseCase) {
	r.Handle("/", createPayment(service),
	).Methods("GET", "OPTIONS").Name("listUsers")
}
