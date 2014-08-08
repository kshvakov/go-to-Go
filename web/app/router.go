package app

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Router struct {
	route map[string]func() string
}

func (r *Router) addRoute(pattern string, handler func() string) {

	r.route[strings.Trim(pattern, "/")] = handler
}

func (r *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	defer func() {

		if panic := recover(); panic != nil {

			log.Printf("panic: %v", panic)

			writer.WriteHeader(http.StatusInternalServerError)

			fmt.Fprint(writer, panic)
		}
	}()
	//fmt.Println(request.URL.Path)
	if f, exist := r.route[strings.Trim(request.URL.Path, "/")]; exist {

		writer.Header().Set("Cache-Control", "no-store, no-cache")

		fmt.Fprint(writer, f())

	} else {

		writer.WriteHeader(http.StatusNotFound)

		fmt.Fprint(writer, "404")
	}
}
