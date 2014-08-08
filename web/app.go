package main

import (
	"github.com/kshvakov/go-to-Go/web/app"
)

func main() {

	a := app.NewApp()

	a.Get("/", func() string { return "Hi!!" })
	a.Get("/a/", func() string { return "url path: a" })
	a.Get("/b/", func() string { return "url path: b" })
	a.Get("/b/a/", func() string { return "url path: b/a" })
	a.Get("/b/:name/", func() string { return "url path: /b/:name/" })
	a.Get("/b/:name/c/", func() string { return "url path: /b/:name/c/" })

	a.Get("/panic/", func() string {

		panic("aaaa!!!!")
		return "panic"
	})

	a.Run()
}
