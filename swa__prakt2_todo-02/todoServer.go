package main

import (
	"net/http"
	"swa__prakt2_todo-02/app/controller"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// URL endpoints for HTML frontend
	r.HandleFunc("/", controller.Index).Methods("GET")
	r.HandleFunc("/login", controller.Login).Methods("POST")
	r.HandleFunc("/logout", controller.Logout).Methods("GET")

	// Static resources
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css", http.FileServer(http.Dir("./static/css"))))
	r.PathPrefix("/fonts/").Handler(http.StripPrefix("/fonts", http.FileServer(http.Dir("./static/fonts"))))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js", http.FileServer(http.Dir("./static/js"))))

	server := http.Server{
		Addr:    ":8443",
		Handler: r,
	}

	server.ListenAndServe()
}
