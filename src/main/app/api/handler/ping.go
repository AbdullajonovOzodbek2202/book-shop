package handler

import "net/http"

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}