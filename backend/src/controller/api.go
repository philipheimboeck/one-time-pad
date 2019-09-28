package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"

	"github.com/gorilla/mux"

	"../dto"
	"../model"
)

// Start the server
func Start(port string, model model.Model) {
	router := mux.NewRouter()
	router.HandleFunc("/api/{key}", makeGetHandler(model)).Methods("GET")
	router.HandleFunc("/api/{key}", makeDeleteHandler(model)).Methods("DELETE")
	router.HandleFunc("/api/", makeStoreHandler(model)).Methods("POST")
	http.ListenAndServe(":"+port, router)
}

func makeGetHandler(model model.Model) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key := vars["key"]
		secret := r.Header.Get("X-Secret")

		v, err := model.Get(key, secret)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json, err := json.Marshal(v)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write(json)
	}
}

func makeDeleteHandler(model model.Model) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key := vars["key"]

		err := model.Delete(key)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func makeStoreHandler(model model.Model) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		secret := r.Header.Get("X-Secret")

		var value dto.ValueDTO
		if err = json.Unmarshal(body, &value); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		uuid, error := uuid.NewRandom()
		if error != nil {
			panic(error)
		}
		key := uuid.String()

		err = model.Store(key, secret, value)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(key))
	}
}
