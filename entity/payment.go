package entity

// PaymentMethod contains payment method info
type PaymentMethod struct {
	URL string
}

// NewPaymentMethod creates new payment method
func NewPaymentMethod(url string) (*PaymentMethod, error) {
	method := &PaymentMethod{
		URL: url,
	}
	// TODO: add validation
	return method, nil
}

