package distance

import (
	"code.google.com/p/gorest"
	distLookup "distance/google"
	"fmt"
	"model"
	"strings"
)

type DistanceService struct {
	gorest.RestService `root:"/ds/" consumes:"application/json" produces:"application/json"`
	getDistance        gorest.EndPoint `method:"GET" path:"/distance/{postcode1:string}/{postcode2:string}" output:"Distance"`
}

func (serv DistanceService) GetDistance(postcode1, postcode2 string) model.Distance {
	fmt.Println("incoming GetDistance request: ", postcode1, postcode2)

	m, err := distLookup.GetDistance(postcode1, postcode2)
	if err != nil {
		if strings.ContainsAny(err.Error(), "Unable to find distance") {		
			serv.ResponseBuilder().SetResponseCode(404).WriteAndOveride([]byte(err.Error()))
		}else{
			serv.ResponseBuilder().SetResponseCode(500).WriteAndOveride([]byte(err.Error()))
		}
	}
	return model.Distance{postcode1, postcode2, m}
}
