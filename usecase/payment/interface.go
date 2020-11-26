package payment

import "github.com/andreybutko/payment/entity"

//UseCase interface
type UseCase interface {
	CreatePayment(productID string) (*entity.PaymentForm, error)
}
