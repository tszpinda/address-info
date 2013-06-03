package main

import (
	address "address"
	"code.google.com/p/gorest"
	distance "distance"
	"net/http"
)

func main() {
	ds := new(distance.DistanceService)
	gorest.RegisterService(ds)
	gorest.RegisterService(new(address.AddressService))
	http.Handle("/", gorest.Handle())
	http.ListenAndServe(":8080", nil)
}
