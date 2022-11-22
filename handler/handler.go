package handler

import (
	"fmt"
	"golangweb/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
		//Biar Tidak Membaca Bawahnya Alias Langsung Return Aja
	}
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))

	if err != nil {
		log.Println(err)
		http.Error(w, "Error Anjir Ntahlah Malas, Err Happening Right Now", http.StatusInternalServerError)
		return
	}

	//Struktur Map[valuekey]value
	// data := map[string]interface{}{
	// 	"title":   "Learning Golang Website",
	// 	"content": "I'm Learning Golang Web With Nico Tamvanzzzzzz",
	// }

	// data := entity.Product{ID: 1, Name: "Honda", Price: 100, Stock: 1}

	data := []entity.Product{
		{ID: 1, Name: "Honda", Price: 100, Stock: 1},
		{ID: 2, Name: "Suzuki", Price: 200, Stock: 8},
		{ID: 3, Name: "Yamaha", Price: 300, Stock: 11},
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error Anjir Ntahlah Malas, Err Happening Right Now", http.StatusInternalServerError)
		return
	}
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World Handler"))
}

func Nico(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Nico Handler"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error Anjir Ntahlah Malas, Err Happening Right Now", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"title":     "Product",
		"idProduct": idNumb,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error Anjir Ntahlah Malas, Err Happening Right Now", http.StatusInternalServerError)
		return
	}
}
