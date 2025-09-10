package main

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/go-chi/chi/v5"
)

type product struct {
	ID   int
	Name string
}

func main() {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query().Get("name")
		if param != "" {
			w.Write([]byte(param))
		} else {
			w.Write([]byte("empty query data"))
		}
	})

	router.Get("/{productName}", func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "productName")
		w.Write([]byte(param))
	})

	router.Get("/json", func(w http.ResponseWriter, r *http.Request) {
		obj := map[string]string{"message": "success"}
		render.JSON(w, r, obj)
	})

	router.Post("/product", func(w http.ResponseWriter, r *http.Request) {
		var product product
		render.DecodeJSON(r.Body, &product)
		product.ID = 5
		render.JSON(w, r, product)
	})
	http.ListenAndServe(":3000", router)
}
