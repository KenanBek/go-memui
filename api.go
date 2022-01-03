package memui

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (mui *MemUI) ServerHTTP() error {
	http.HandleFunc("/", mui.handleIndex)
	http.HandleFunc("/explore", mui.handleExplore)

	return http.ListenAndServe(":8080", nil)
}

func (mui *MemUI) handleIndex(w http.ResponseWriter, r *http.Request) {
	if len(mui.memObjects) == 0 {
		w.Write([]byte("No objects registered"))
		return
	}

	w.Write([]byte("<html><body><ul>"))
	for objType := range mui.memObjects {
		w.Write([]byte(fmt.Sprintf("<li><a href=\"/explore?type=%s\">%s</a></li>", objType, objType)))
	}
	w.Write([]byte("</ul></body></html>"))
}

func (mui *MemUI) handleExplore(w http.ResponseWriter, r *http.Request) {
	objType := r.URL.Query().Get("type")
	if objType == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	objList, ok := mui.memObjects[objType]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	respBody, err := json.Marshal(objList)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}
