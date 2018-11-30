package main

import (
	"sync"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type Router interface {
	Init() *chi.Mux
}

type router struct{}

func (router *router) Init() *chi.Mux {

	duler := Depedency().InjectDuler()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.Timeout(60 * time.Second))

	// route
	r.Get("/duler/{id}", duler.Find)
	r.Get("/dulers", duler.FindAll)
	r.Post("/store", duler.Store)

	return r
}

func Route() Router {

	var (
		r    *router
		once sync.Once
	)

	if r == nil {
		once.Do(func() {
			r = &router{}
		})
	}
	return r
}
