package api

import (
	"errors"
	"keystra/api/endpoint"
	"keystra/api/middleware"
	"keystra/api/types"
	"keystra/storage/sqlite"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

const (
	port          = ":4000"
	database_name = "keystra.db"
)

type KeystraServer struct {
	keystraAPI *types.KeystraAPI
}

type API interface {
	Close() error
	Start() error
}

func NewApi() (*KeystraServer, error) {
	database, err := sqlite.Connect(database_name)

	if err != nil {
		return nil, err
	}

	router := chi.NewRouter()
	srv := &http.Server{
		Addr:         port,
		Handler:      router,
		ErrorLog:     log.Default(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	router.Use(middleware.Default(database)...)

	kAPI := &types.KeystraAPI{
		Router:  router,
		Storage: database,
		Server:  srv,
	}

	endpoint.RegisterRoutes(kAPI)

	return &KeystraServer{keystraAPI: kAPI}, nil
}

func (ks *KeystraServer) Close() error {
	var err error

	log.Println("Closing storage")
	err = ks.keystraAPI.Storage.Close()

	if err != nil {
		return errors.New("error closing database: " + err.Error())
	}

	log.Println("Closing server")
	err = ks.keystraAPI.Server.Close()

	if err != nil {
		return errors.New("error shutting down server: " + err.Error())
	}

	return nil
}

func (ks *KeystraServer) Start() error {
	log.Println("starting server to listen on port:", port)

	err := ks.keystraAPI.Server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
