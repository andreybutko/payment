package payment

import "github.com/andreybutko/payment/entity"

//UseCase interface
type UseCase interface {
	GetPaymentForm(productID string) (*entity.PaymentForm, error)
}
