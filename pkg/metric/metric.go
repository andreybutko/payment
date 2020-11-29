package metric

import (
	"log"
)

//service implements Service interface
type service struct {
}

func NewMetricService() (*service, error) {
	s := &service{}

	return s, nil
}

//SaveHTTP send metrics to server
func (s *service) SaveHTTP(h *HTTP) {
	log.Printf("Method %s executed\n Duration: %f", h.Method, h.Duration)
}