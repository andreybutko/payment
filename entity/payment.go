package entity

// PaymentForm contains form info
type PaymentForm struct {
	URL string
}
// NewPaymentForm creates new payment form
func NewPaymentForm(url string) (*PaymentForm, error) {
	form := &PaymentForm{
		URL: url,
	}
	// TODO: add validation
	return form, nil
}

