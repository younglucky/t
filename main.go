package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"log"
)

var (
	router *mux.Router
	r      *render.Render
)

func init() {
	log.Println("main.go init()...")

	router = mux.NewRouter()

	r = render.New(render.Options{
		Delims:        render.Delims{"[[", "]]"},
		IsDevelopment: true,
	})
}

func main() {
	log.Println("main.go main()...")

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":3000")
}
