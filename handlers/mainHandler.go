package handlers

import (
	"fmt"
	"net/http"
)

var MainHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	payload := fmt.Sprintf("Hello %d", http.StatusOK)
	_, _ = w.Write([]byte(payload))

})
