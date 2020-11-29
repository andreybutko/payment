package payment

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_GetPaymentMethods(t *testing.T) {
	// Replace with mock from testify
	provider := newProviderMock()
	service := NewService(provider)
	res, err := service.GetPaymentMethods("productID")
	
	assert.Nil(t, err, "Error should not be encountered")
	assert.NotNil(t, res, "Response should not be empty")
}
