package middleware

import (
	"database/sql"
	"net/http"

	chiMid "github.com/go-chi/chi/v5/middleware"
)

func Default(database *sql.DB) []func(http.Handler) http.Handler {
	return []func(http.Handler) http.Handler{
		responseTypeJSON,
		databaseHeartbeat(database),
		chiMid.Heartbeat("/ping"),
		chiMid.RequestID,
		chiMid.Logger,
		chiMid.Recoverer,
	}

}
