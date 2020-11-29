package payment

import (
	"github.com/andreybutko/payment/entity"
	"log"
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

// GetPaymentMethods returns payment methods from provider
func (s *Service) GetPaymentMethods(productID string) (*[]entity.PaymentMethod, error) {
	log.Printf("ProductID: %s - processed", productID)
	payment, err := s.provider.getPayment(&Request{})
	if err != nil {
		return nil, err
	}
	method, _ := entity.NewPaymentMethod(payment.URL)

	return &[]entity.PaymentMethod{*method}, nil
}
