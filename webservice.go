package main

import "net/http"

// WebService is the interface that should be implemented by types that want to
// provide web services.
type WebService interface {
	// GetPath returns the path to be associated with the service.
	GetPath() string
	WebPosthandler(w http.ResponseWriter, r *http.Request) (int, string)
}
