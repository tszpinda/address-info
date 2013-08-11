package main

import (
	"address"
	"code.google.com/p/gorest"
	"distance"
	"net/http"
	"web"
)

func serveSingle(pattern string, filename string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}

func main() {
	ds := new(distance.DistanceService)
	gorest.RegisterService(ds)
	gorest.RegisterService(new(address.AddressService))

	serveSingle("/favicon.ico", "./favicon.ico")
	view.Mount()

	http.Handle("/", gorest.Handle())
	http.ListenAndServe(":8080", nil)
}
