package main

import (
	"github.com/gorilla/mux"
	"github.com/rhass99/wixoftrials/api"
	"net/http"
)

var r = mux.NewRouter()


func main() {

	//r.HandleFunc("/api/user/profile", api.HandleProfile)
	r.HandleFunc("/api/guest/signup", api.HandleAccountSignup).Methods("POST")
	r.HandleFunc("/api/guest/getall", api.HandleRetrieveAll).Methods("GET")
	//r.HandleFunc("/api/guest/login", api.HandleLogin)
	//r.HandleFunc("/api/guest/", api.HandleIndex)

	http.ListenAndServe(":3000", r)
}

