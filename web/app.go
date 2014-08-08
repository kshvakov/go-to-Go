package main

import (
	"github.com/kshvakov/go-to-Go/web/app"
)

func main() {

	a := app.NewApp()

	a.Get("/", func() string { return "Hi!!" })
	a.Get("/a/", func() string { return "url path: a" })
	a.Get("/b/", func() string { return "url path: b" })

	a.Get("/panic/", func() string {

		panic("aaaa!!!!")
		return "panic"
	})

	a.Start()
}
