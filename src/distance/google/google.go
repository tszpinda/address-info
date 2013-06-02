package google

import (
	"filecache"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"encoding/json"
)

//used to Unmarshal response from google, requires public access of the fields (so needs to be upercase) 
type DistanceResponse struct {
	Routes []struct {
		Legs []struct {
			Distance struct {
				Text  string
				Value float64
			}
		}
	}
}
func (dr DistanceResponse) getDistance() (meters float64) {
	distance := dr.Routes[0].Legs[0].Distance
	meters = distance.Value
	return
}
func GetDistance(p1, p2 string) (meters float64) {
	p1 = url.QueryEscape(p1)
	p2 = url.QueryEscape(p2)
	d, f := filecache.GetDistance(p1, p2)
	if f {
		return d.Meters
	}

	driveTime := "http://maps.googleapis.com/maps/api/directions/json?origin=%s&destination=%s&sensor=false"
	driveTimeUrl := fmt.Sprintf(driveTime, p1, p2)
	fmt.Println(driveTimeUrl)

	res, err := http.Get(driveTimeUrl)
	if err != nil {
		log.Fatal(err)
	}
	rawResponse, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		//TODO cope with error
		log.Fatal(err)
	}

	distanceResponse := DistanceResponse{}
	if err := json.Unmarshal(rawResponse, &distanceResponse); err == nil {
		meters = distanceResponse.getDistance()
		filecache.CacheDistance(p1, p2, meters)
	} else {
		//TODO cope with error
		log.Fatal(err)
	}

	return
}