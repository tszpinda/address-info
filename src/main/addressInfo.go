package main

import (
	addrLookup "addrLookup/google"
	distLookup "distance/google"
	"fmt"
)

func main() {
	postcode := "EX16 6AB"
	address, err := addrLookup.GetAddress(postcode)
	if err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		fmt.Printf("%+v\n", address)
	}
	
	fmt.Println("###### Check distance #####")
	
	postcode2 := "EX16 4QA"
	m, err := distLookup.GetDistance(postcode, postcode2)
	if err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		fmt.Printf("q1: %v \n", m)
	}
	m, err = distLookup.GetDistance(postcode2, postcode)
	if err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		fmt.Printf("q2: %v \n", m)
	}
}
