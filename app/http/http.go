package http

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Service satisfies the http Handler interface.
type Service struct{}

// New returns a new Service.
func New() *Service {
	return &Service{}
}

// ServeHTTP implements the interface method.
func (s *Service) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	log.Printf("received request from: %s", r.RemoteAddr)

	// Apply some random delay.
	n := rand.Intn(10) // n will be between 0 and 10
	time.Sleep(time.Duration(n * 50) * time.Millisecond)

	// Just a greeting.
	if _, err := rw.Write([]byte("Greetings!")); err != nil {
		log.Printf("writing response failed: %v", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}
