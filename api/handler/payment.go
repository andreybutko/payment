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

func getPaymentMethods(service payment.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Extract error message from method
		errorMessage := "Error getting payment methods."
		vars := mux.Vars(r)
		id := vars["id"]

		methods, err := service.GetPaymentMethods(id)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}

		if methods == nil {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
		var response []presenter.PaymentMethod
		for _, v := range *methods {
			response = append(response, presenter.PaymentMethod{
				URL: v.URL,
			})
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
		}
	})
}

//MakePaymentHandlers creates url handlers
func MakePaymentHandlers(r *mux.Router, n negroni.Negroni, service payment.UseCase) {
	r.Handle("/products/{id}/payment_methods", n.With(
		negroni.Wrap(getPaymentMethods(service))),
	).Methods("GET", "OPTIONS").Name("getPaymentMethods")
}
