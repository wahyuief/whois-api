package handler

import (
	"fmt"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/someshkar/whois-api/lib"
	"github.com/someshkar/whois-api/structs"
)

// MainHandler handles Whois info for a single domain
func MainHandler(w http.ResponseWriter, r *http.Request) {

	// Make sure it's a POST request
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Please only POST data here")
		return
	}

	// Decode JSON body
	decoder := jsoniter.NewDecoder(r.Body)
	var body structs.SingleBody

	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get Whois data
	whois, err := lib.GetWhois(body.Domain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert Whois data to JSON
	jsonWhois, err := jsoniter.Marshal(whois)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonWhois)
}
