package main

import (
	"golangweb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// mux.HandleFunc("/about", aboutHandler)
	// aboutHandler := func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("About Page Handler"))
	// }

	// mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Profile Page Handler"))
	// })

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/hello", handler.HelloWorld)
	mux.HandleFunc("/nico", handler.Nico)
	mux.HandleFunc("/product", handler.ProductHandler)

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting On Port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
