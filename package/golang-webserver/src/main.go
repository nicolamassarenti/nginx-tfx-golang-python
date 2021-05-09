package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Response struct {
	lemmas string
	sentiment float64
}

func handler(w http.ResponseWriter, r *http.Request) {

	sentences, ok := r.URL.Query()["sentence"]

	if !ok || len(sentences[0]) < 1 {
		log.Println("Url Param 'sentence' is missing")
		return
	}

	sentence := sentences[0]

	const URL = "http://127.0.0.1:8081/api/nlp"
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("sentence", sentence)
	req.URL.RawQuery = q.Encode()

	var getResponse Response
	get, err := http.Get(req.URL.String())
	if err != nil {
		return
	}
	err = json.NewDecoder(get.Body).Decode(&getResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var body []byte
	if body, err = json.Marshal(getResponse); err != nil {
		return
	}

	w.WriteHeader(200)
	w.Write(body)

	return
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}