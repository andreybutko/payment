package restclient

import (
	"net/http"
	"github.com/andreybutko/payment/api/clients/mock"
)

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &mock.MockClient{}
}

// Get sends a Get request to the URL
func Get(url string, headers http.Header) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header = headers
	return Client.Do(request)
}