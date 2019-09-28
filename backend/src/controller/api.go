package controller

import "net/http"

// Start the server
func Start() {
	http.HandleFunc("/api", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
