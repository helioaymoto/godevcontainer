package impl

import (
	"encoding/json"
	"net/http"
)

func TestGet(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World!")
}
