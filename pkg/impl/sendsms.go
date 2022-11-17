package impl

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type User struct {
	App     string `json:"App"`
	Source  string `json:"Source"`
	Misdn   string `json:"Misdn"`
	Message string `json:"Message"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	b, _ := json.Marshal(user)

	if strings.Contains(user.Message, "REAUTHORIZATION") {
		fmt.Println("Sending Request to SG to release the Redirection...")
	}
	log.Printf("POST Received: %s\n", string(b))
	json.NewEncoder(w).Encode(user)
}
