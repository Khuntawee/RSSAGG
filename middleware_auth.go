package main

import (
	"net/http"

	"github.com/WhyURead/rssagg/internal/auth"
	"github.com/WhyURead/rssagg/internal/database"
)


type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondError(w, 403, "Cannot get API key")
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondError(w, 400, "Cannot get user")
			return
		}

		handler(w, r, user)
	}
}