package main

import (
	//"fmt"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Frutas struct {
	Id       int
	Fruta    string
	Cantidad int
}

var frutas = []Frutas{
	{Id: 1, Fruta: "Manzana", Cantidad: 7},
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
	json.NewEncoder(w).Encode(frutas)
}

func postData(w http.ResponseWriter, r *http.Request) {
	var data Frutas
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	busqueda := 5
	existe := existeEnArreglo(frutas, busqueda)
	fmt.Println("Buscando %d en arreglo. ¿Existe? %t", busqueda, existe)

	frutas = append(frutas, data)
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(data)
}

func existeEnArreglo(arreglo []int, busqueda int) bool {
	for _, numero := range arreglo {
		if numero == busqueda {
			return true
		}
	}
	return false
}
