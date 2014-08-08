package app

import (
	"log"
	"net/http"
	"time"
)

func NewApp() *App {
	return &App{
		router: &Router{route: make(map[string]func() string)},
	}
}

type App struct {
	router *Router
}

func (a *App) Get(pattern string, handler func() string) {

	a.router.addRoute(pattern, handler)
}

func (a *App) Run() {

	httpServer := &http.Server{
		Addr:         ":8081",
		Handler:      a.router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Print("Start")
	log.Fatal(httpServer.ListenAndServe())
}
