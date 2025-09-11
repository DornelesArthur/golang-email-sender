package main

import (
	"golang-email-sender/internal/contract"
	"golang-email-sender/internal/domain/campaign"
	"net/http"

	"github.com/go-chi/render"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	service := campaign.Service{}

	router.Post("/campaigns", func(w http.ResponseWriter, r *http.Request) {
		var campaign contract.NewCampaign
		err := render.DecodeJSON(r.Body, &campaign)
		if err != nil {
			render.Status(r, 500)
			render.JSON(w, r, map[string]string{"error": "Internal Server Error"})
			return
		}
		id, err := service.Create(campaign)

		if err != nil {
			render.Status(r, 400)
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
		render.Status(r, 201)
		render.JSON(w, r, map[string]string{"id": id})
	})

	http.ListenAndServe(":3000", router)
}
