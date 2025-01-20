package endpoint

import (
	"encoding/json"
	"keystra/api/types"
	"keystra/service"
	"net/http"
)

func RegisterRoutes(k *types.KeystraAPI) {
	k.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Welcome to Keystra project")
	})

	services := service.NewService(k.Storage)

	registerUserHandlers(k.Router, services.User)
}
