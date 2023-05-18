package main

import (
	//"fmt"
	"encoding/json"
	"log"
	"net/http"
)

type Persona struct {
	Id        int
	Nombre    string
	Direccion string
}

var personas = []Persona{
	{Id: 1, Nombre: "Gerardo", Direccion: "CDMX"},
}

func main() {

	http.HandleFunc("/data", dataHan)
	log.Fatal(http.ListenAndServe("8080", nil))
}

func dataHan(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getData(w, r)
	case "POST":
		postData(w, r)
	default:
		http.Error(w, "No existe este método", http.StatusMethodNotAllowed)
	}
}

func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(personas)
}

func postData(w http.ResponseWriter, r *http.Request) {
	var data Persona
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	personas = append(personas, data)
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(data)
}
