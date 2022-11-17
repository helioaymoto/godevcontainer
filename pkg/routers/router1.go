package routers

import (
	"log"
	"net/http"

	"github.com/helioaymoto/training/pkg/impl"

	"github.com/gorilla/mux"
)

func InitializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", impl.TestGet).Methods("GET")
	log.Fatal(http.ListenAndServe(":9000", r))

}
