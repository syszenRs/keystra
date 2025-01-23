package endpoint

import (
	"encoding/json"
	"keystra/internal/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type resource struct {
	service service.UserService
}

func registerUserHandlers(router *chi.Mux, service service.UserService) {
	s := &resource{
		service: service,
	}

	router.Route("/user/{id}", func(r chi.Router) {
		r.Get("/", s.get)
		r.Delete("/", s.delete)
	})
}

func (res *resource) get(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	if id == 0 {
		http.Error(w, "missing user identifier", http.StatusBadRequest)
	}

	user, err := res.service.Get(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(user)
}

func (res *resource) delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	if id == 0 {
		http.Error(w, "missing user identifier", http.StatusBadRequest)
	}

	err := res.service.Delete(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
