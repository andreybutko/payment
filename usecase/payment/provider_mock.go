package payment

type mockProvider struct {
}

func newProviderMock() *mockProvider {
	return &mockProvider{}
}

func (p *mockProvider) getPayment(req *Request) (*Response, error) {
	return &Response{URL: "test.com"}, nil
}
