package memui

import (
	"net/http"
)

func (mui *MemUI) ServerHTTP() error {
	http.HandleFunc("/", mui.handleIndex)

	return http.ListenAndServe(":8080", nil)
}

func (mui *MemUI) handleIndex(w http.ResponseWriter, r *http.Request) {
	if len(mui.memObjects) == 0 {
		w.Write([]byte("No objects registered"))
		return
	}

	w.Write([]byte("<html><body><ul>"))
	for objType := range mui.memObjects {
		w.Write([]byte("<li>" + objType + "</li>"))
	}
	w.Write([]byte("</ul></body></html>"))
}
