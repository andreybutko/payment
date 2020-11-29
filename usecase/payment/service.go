package payment

import (
	"github.com/andreybutko/payment/entity"
)

// Service manages payment process
type Service struct {
	provider Provider
}

// NewService returns new payment service
func NewService(provider Provider) *Service {
	return &Service{
		provider: provider,
	}
}

// GetPaymentForm returns payment form from provider
func (s *Service) GetPaymentForm(productID string) (*entity.PaymentForm, error) {
	payment, err := s.provider.getPayment(&Request{})
	if err != nil {
		return nil, err
	}
	form, _ := entity.NewPaymentForm(payment.URL)

	return form, nil
}
