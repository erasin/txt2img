package main

import

//"github.com/facebookgo/grace/gracehttp"
"github.com/lunny/tango"

func main() {
	tg := tango.Classic()

	//tg.Get("/:to/:from/:date", new(TxtImgCon))
	tg.Get("/", new(TxtImgCon))

	tg.Run(":9092")
	//err := gracehttp.Serve(&http.Server{Addr: ":9092", Handler: tg})
	//if err != nil {
	//	log.Fatalln(err)
	//}
}
