package types

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type KeystraAPI struct {
	Router  *chi.Mux
	Storage *sql.DB
	Server  *http.Server
}
