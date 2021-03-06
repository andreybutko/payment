package payment

import "github.com/andreybutko/payment/entity"

//UseCase interface
type UseCase interface {
	GetPaymentMethods(productID string) (*[]entity.PaymentMethod, error)
}

// Request contains payment response fields
type Request struct {
}

// Response contains payment request fields
type Response struct {
	URL string
}

// Provider provides payment operations
type Provider interface {
	getPayment(req *Request)  (*Response, error)
}
