package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

//"github.com/facebookgo/grace/gracehttp"

func main() {

	//tg.Get("/:to/:from/:date", new(TxtImgCon))
	//tg.Get("/", new(TxtImgCon))

	r := mux.NewRouter()
	r.HandleFunc("/", TxtImgConGet)

	tg := negroni.Classic()
	tg.UseHandler(r)
	tg.Run(":9092")
	//err := gracehttp.Serve(&http.Server{Addr: ":9092", Handler: tg})
	//if err != nil {
	//	log.Fatalln(err)
	//}
}
