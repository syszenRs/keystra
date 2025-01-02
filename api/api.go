package api

import (
	"encoding/json"
	"keystone/api/prep"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const PORT = ":4000" //TODO: port as environment variable

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Welcome to keystone project")
}

func test(w http.ResponseWriter, r *http.Request) {
	res := prep.RemoveKAdjacent("deeedbbcccbdaa", 3)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(res)
}

func setup(router *chi.Mux) {
	router.Use(middleware.Heartbeat("/ping"))
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", home)
	router.Get("/test", test)
}

func Start() {
	router := chi.NewRouter()

	setup(router)

	srv := http.Server{
		Addr:     PORT,
		Handler:  router,
		ErrorLog: log.Default(),
	}

	log.Println("starting server to listen on port:", PORT)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
