package main

import (
	"./app"
)

func main() {

	a := app.NewApp()

	a.AddRoute("/", func() string { return "Hi!!" })
	a.AddRoute("/a/", func() string { return "url path: a" })
	a.AddRoute("/b/", func() string { return "url path: b" })

	a.AddRoute("/panic/", func() string {

		panic("aaaa!!!!")
		return "panic"
	})

	a.Start()
}
