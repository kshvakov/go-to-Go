package app

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func NewApp() *App {
	return &App{
		route: make(map[string]func() string),
	}
}

type App struct {
	route map[string]func() string
}

func (a *App) Get(pattern string, handler func() string) {

	a.route[strings.Trim(pattern, "/")] = handler
}

func (a *App) handle(writer http.ResponseWriter, request *http.Request) {

	defer func() {

		if panic := recover(); panic != nil {

			log.Printf("panic: %v", panic)

			writer.WriteHeader(http.StatusInternalServerError)

			fmt.Fprint(writer, panic)
		}
	}()

	if f, exist := a.route[strings.Trim(request.URL.Path, "/")]; exist {

		writer.Header().Set("Cache-Control", "no-store, no-cache")

		fmt.Fprint(writer, f())

	} else {

		writer.WriteHeader(http.StatusNotFound)

		fmt.Fprint(writer, "404")
	}
}

func (a *App) Start() {

	httpServer := &http.Server{
		Addr:         ":8081",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	http.HandleFunc("/", a.handle)

	log.Print("Start")
	log.Fatal(httpServer.ListenAndServe())
}
