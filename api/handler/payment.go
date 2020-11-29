package handler

import (
	"github.com/codegangsta/negroni"
	"github.com/andreybutko/payment/api/presenter"
	"encoding/json"
	"github.com/andreybutko/payment/usecase/payment"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func getPaymentForm(service payment.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error getting payment form."
		vars := mux.Vars(r)
		id := vars["productID"]

		data, err := service.GetPaymentForm(id)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}

		form := &presenter.PaymentForm{
			URL: data.URL,
		}

		if err := json.NewEncoder(w).Encode(form); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

//MakePaymentHandlers creates url handlers
func MakePaymentHandlers(r *mux.Router, n negroni.Negroni, service payment.UseCase) {
	r.Handle("/payments/{productID}", n.With(
		negroni.Wrap(getPaymentForm(service))),
	).Methods("GET", "OPTIONS").Name("getPaymentForm")
}
