package main

import (
	"log"
	"net/http"

	"github.com/baekdavid/FriGO_chef/app"
	"github.com/urfave/negroni"
)

func main() {
	m := app.MakeNewHandler()
	n := negroni.Classic()
	n.UseHandler()

	log.Println("Started App")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}
