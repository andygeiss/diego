package handlers

import (
	"net/http"
)

// StaticDataHandler ...
type StaticDataHandler struct {
	data []byte
	mime string
}

// NewStaticDataHandler ...
func NewStaticDataHandler(data []byte, mime string) http.Handler {
	return &StaticDataHandler{
		data: data,
		mime: mime,
	}
}

// ServeHTTP serves the current content to the client.
func (h *StaticDataHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", h.mime)
	res.Write(h.data)
}
