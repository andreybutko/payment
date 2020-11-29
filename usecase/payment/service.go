package payment

import (
	"encoding/json"
	"github.com/andreybutko/payment/api/clients/mock"
	"github.com/andreybutko/payment/api/clients/restclient"
	"github.com/andreybutko/payment/entity"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetPaymentForm(productID string) (*entity.PaymentForm, error) {

	// TODO: move outside usecase logic
	obj := `{"url":"payment.com/pay"}`
	mock.HTTPResponse(obj, 200)

	res, _ := restclient.Get("example.com", nil)
	form := entity.PaymentForm{}
	err := json.NewDecoder(res.Body).Decode(&form)

	if err != nil {
		return nil, err
	}
	return &form, nil
}
