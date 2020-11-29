package payment

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_GetPaymentForm(t *testing.T) {
	// Replace with mock from testify
	provider := newProviderMock()
	service := NewService(provider)
	res, err := service.GetPaymentForm("productID")
	
	assert.Nil(t, err, "Error should not be encountered")
	assert.NotNil(t, res, "Response should not be empty")
	assert.Equal(t, res.URL, "test.com", "Response should correspond to expected")
}
