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

func getVagas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content=Type", "aplication/json")

	json.NewEncoder(w).Encode(request())

}
