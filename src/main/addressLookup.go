package main

import (
	"fmt"
	"net/url"
//	"model"
	distLookup "distance/google"
	addrLookup "addrLookup/google"
)


func main() {
	postcode := url.QueryEscape("EX16 6AB")
	address := addrLookup.GetAddress(postcode)
	fmt.Printf("%+v\n", address)

	postcode2 := url.QueryEscape("EX16 4QA")
	m := distLookup.GetDistance(postcode, postcode2)
	fmt.Printf("q1: %v \n", m)
	m = distLookup.GetDistance(postcode2, postcode)
	fmt.Printf("q2: %v \n", m)
}
