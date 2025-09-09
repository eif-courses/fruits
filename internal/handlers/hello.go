package handlers

import "net/http"

func HelloResponse(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World!"))
	if err != nil {
		return
	}
}
