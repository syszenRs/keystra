package middleware

import (
	"database/sql"
	"net/http"
)

func databaseHeartbeat(database *sql.DB) func(http.Handler) http.Handler {
	fn := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := database.Ping()
			if err != nil {
				http.Error(w, "database not available."+err.Error(), http.StatusServiceUnavailable)
			} else {
				next.ServeHTTP(w, r)
			}
		})
	}

	return fn
}
