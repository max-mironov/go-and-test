package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/apple-app-site-association", AASAHandler)
	r.HandleFunc("/applinks", AppLinksHandler)
	r.HandleFunc("/fbtest", FBMessengerHandler)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	http.ServeFile(w, r, "html/index.html")
}

func AppLinksHandler(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Prefer-Html-Meta-Tags", "al")
	w.WriteHeader(http.StatusOK)
	http.ServeFile(w, r, "html/applink.html")
}

func FBMessengerHandler(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Prefer-Html-Meta-Tags", "al")
	w.WriteHeader(http.StatusOK)
	http.ServeFile(w, r, "html/fbtest.html")
}

func AASAHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"applinks":{"apps":[],"details":[{"appID":"285W24646W.com.akvelon.adjust","paths":["/*"]}]}}`)
}
