package test_learning

import "net/http"

func httpHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
