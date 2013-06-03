package addrLookup

import (
	addrLookup "addrLookup/google"
	"code.google.com/p/gorest"
	"fmt"
	"model"
	"strings"
)

type AddressService struct {
	gorest.RestService `root:"/ds/" consumes:"application/json" produces:"application/json"`
	getAddress         gorest.EndPoint `method:"GET" path:"/address/{postcode:string}" output:"Address"`
}

func (serv AddressService) GetAddress(postcode string) model.Address {
	fmt.Println("incoming GetAddress request: ", postcode)

	a, err := addrLookup.GetAddress(postcode)
	if err != nil {
		if strings.ContainsAny(err.Error(), "Unable to find address") {		
			serv.ResponseBuilder().SetResponseCode(404).WriteAndOveride([]byte("Postcode '" + postcode + "' not found"))
		}else{
			serv.ResponseBuilder().SetResponseCode(500).WriteAndOveride([]byte(err.Error()))
		}
	}
	return a
}
