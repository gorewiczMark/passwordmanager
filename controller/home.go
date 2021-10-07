package controller

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type home struct {
	homeTemplate *template.Template
}

func (h *home) registerRoutes(r *mux.Router) {

	r.HandleFunc("/home", h.handleHome)
	r.HandleFunc("/", h.handleHome)
}

func (h *home) handleHome(w http.ResponseWriter, r *http.Request) {
	err := h.homeTemplate.ExecuteTemplate(w, "_layout.html", passEntries)
	if err != nil {
		log.Fatal(err)
	}

}
