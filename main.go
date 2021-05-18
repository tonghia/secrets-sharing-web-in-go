package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

func secretHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		createSecret(w, r)
	case "GET":
		getSecret(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func healthcheck(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "healthcheck")
}

func main() {
	mux := http.NewServeMux()

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	mux.HandleFunc("/healthcheck", healthcheck)
	mux.HandleFunc("/", secretHandler)

	server.ListenAndServe()
}

var secretStore = memoryStore{Mu: sync.Mutex{}, Store: make(map[string]string)}

func getHash(plainText string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(plainText)))
}

func createSecret(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading body", http.StatusInternalServerError)
		return
	}
	p := createSecretPayload{}
	err = json.Unmarshal(bytes, &p)
	if err != nil || len(p.PlainText) == 0 {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	digest := getHash(p.PlainText)
	resp := createSecretResponse{Id: digest}

	s := secretData{Id: resp.Id, Secret: p.PlainText}
	secretStore.Write(s)
	jd, err := json.Marshal(&resp)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jd)
}

type getSecretResponse struct {
	Data string `json:"data"`
}

func getSecret(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path
	id = strings.TrimPrefix(id, "/")
	resp := getSecretResponse{}
	v := secretStore.Read(id)
	resp.Data = v
	jd, err := json.Marshal(&resp)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	if len(resp.Data) == 0 {
		w.WriteHeader(404)
	}
	w.Write(jd)
}
