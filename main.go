package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/IndBrony/lettercounter"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type response struct {
	Vocals     int `json:"vocals"`
	Consonants int `json:"consonants"`
}

func handler(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	text := request.Form.Get("text")
	vocals, consonants := lettercounter.NumberOfVocalsAndConsonantsConcurrent(text)

	response := response{
		Vocals:     vocals,
		Consonants: consonants,
	}

	writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		log.Printf("Error writing json: %v", err)
		http.Error(writer, "can't write json", http.StatusInternalServerError)
		return
	}
}
