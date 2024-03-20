package testgrp

import (
	"encoding/json"
	"net/http"
)

// Test is our example route
func Test(w http.ResponseWriter, r *http.Request) {
	// validate request data
	// business layer
	// return errors
	// handle OK response

	status := struct {
		Status string
	}{
		Status: "OK",
	}

	json.NewEncoder(w).Encode(status)
}
