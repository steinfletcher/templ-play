package web

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/steinfletcher/templ-play/web/layout"
	"github.com/steinfletcher/templ-play/web/todos"
)

func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/todos", getTodos)
	return r
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	err := layout.IndexPage(todos.Get()).Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("something went wrong"))
	}
}
