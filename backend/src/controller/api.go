package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"../domain"
	"../model"
)

// Start the server
func Start(model model.Model) {
	router := mux.NewRouter()
	router.HandleFunc("/api/{key}", makeGetHandler(model)).Methods("GET")
	router.HandleFunc("/api/{key}", makeDeleteHandler(model)).Methods("DELETE")
	router.HandleFunc("/api/", makeStoreHandler(model)).Methods("POST")
	http.ListenAndServe(":8080", router)
}

func makeGetHandler(model model.Model) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key := vars["key"]

		v := model.Get(key)

		json, _ := json.Marshal(v)
		w.Write(json)
	}
}

func makeDeleteHandler(model model.Model) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key := vars["key"]

		model.Delete(key)

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

		var value domain.Value
		if err = json.Unmarshal(body, &value); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		model.Store(value)

		w.WriteHeader(http.StatusCreated)
	}
}
