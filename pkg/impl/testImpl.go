package impl

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func TestGet(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("Cannot get hostname")
		os.Exit(1)
	}

	json.NewEncoder(w).Encode("Hello World from " + hostname)
}
