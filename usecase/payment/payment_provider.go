package payment

import (
	"encoding/json"
	"github.com/andreybutko/payment/api/clients/mock"
	"github.com/andreybutko/payment/api/clients/restclient"
)

// ServiceProvider providers payment operations
type ServiceProvider struct {
}

// NewProvider returns new payment provider
func NewProvider() *ServiceProvider {
	return &ServiceProvider {}
}

func (p *ServiceProvider) getPayment(req *Request) (*Response, error) {
	obj := `{"url":"payment.com/pay"}`
	mock.HTTPResponse(obj, 200)

	res, _ := restclient.Get("example.com", nil)
	payment := Response{}
	err := json.NewDecoder(res.Body).Decode(&payment)

	if err != nil {
		return nil, err
	}

	return &payment, nil
}