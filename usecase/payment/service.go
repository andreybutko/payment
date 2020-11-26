package payment

import "github.com/andreybutko/payment/entity"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreatePayment(productID string) (*entity.PaymentForm, error) {
	// TODO: Replace with URL from service
	form, err := entity.NewPaymentForm("url")
	if err == nil {
		return nil, err
	}

	return form, nil
}
