package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/apple-app-site-association", AASAHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	http.ServeFile(w, r, "html/index.html")
}

func AASAHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"applinks":{"apps":[],"details":[{"appID":"285W24646W.com.akvelonadjust.examples.com","paths":["/*"]}]}}`)
}
