package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/vagas", getVagas)
	fmt.Println("api in on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

type Vagas struct {
	Hash float64
	Name string
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func getVagas(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	w.Header().Set("Content=Type", "aplication/json")

	interships := request()
	for _, i := range interships {
		json.NewEncoder(w).Encode(i)
	}
}
