package google

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"model"
	"net/url"
)


type PostCodeResponse struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Lat float32
				Lng float32
			}
		}
	}
}

type AddressResponse struct {
	Results []struct {
		Address_components []struct {
			Long_name string
			Types     []string
		}
	}
}


func getGeometry(postcode string) (lat, lng float32) {
	postcode = url.QueryEscape(postcode)
	urlGeometry := "http://maps.googleapis.com/maps/api/geocode/json?address=" + postcode + "&sensor=false"
	res, err := http.Get(urlGeometry)
	if err != nil {
		log.Fatal(err)
	}
	rawResponse, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		//TODO cope with error
		log.Fatal(err)
	}

	postCodeJson := PostCodeResponse{}
	if err := json.Unmarshal(rawResponse, &postCodeJson); err == nil {
		//TODO cope with no results
		lat = postCodeJson.Results[0].Geometry.Location.Lat
		lng = postCodeJson.Results[0].Geometry.Location.Lng
	} else {
		//TODO cope with error
		log.Fatal(err)
	}
	return
}

func contains(sl []string, text string) bool {
	for _, v := range sl {
		if strings.EqualFold(v, text) {
			return true
		}
	}
	return false
}

func GetAddress(postcode string) model.Address {
	lat, lng := getGeometry(postcode)
	address := getAddress(lat, lng)
	return address
}
func getAddress(lat, lng float32) model.Address {
	fmt.Println("lat: %g lng: %g", lat, lng)
	urlTemplate := "http://maps.googleapis.com/maps/api/geocode/json?latlng=%g,%g&sensor=false"
	geocodeUrl := fmt.Sprintf(urlTemplate, lat, lng)
	fmt.Println(geocodeUrl)

	res, err := http.Get(geocodeUrl)
	if err != nil {
		log.Fatal(err)
	}
	rawResponse, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		//TODO cope with error
		log.Fatal(err)
	}

	addressResponse := AddressResponse{}
	address := model.Address{}
	if err := json.Unmarshal(rawResponse, &addressResponse); err == nil {
		//get first one as its most accurate
		fmt.Println(addressResponse.Results[0])
		//TODO cope with no results
		//longName := addressResponse.Results[0].address_components[0].long_name

		addressSlice := addressResponse.Results[0]
		addressData := addressSlice.Address_components
		for i := 0; i < len(addressData); i++ {
			addressElem := addressData[i]
			//if addressElem.Long_name
			fmt.Printf("p[%v] == %v\n", i, addressElem.Long_name)
			if contains(addressElem.Types, "street_number") {
				address.HouseNumber = addressElem.Long_name
			} else if contains(addressElem.Types, "route") {
				address.Street = addressElem.Long_name
			} else if contains(addressElem.Types, "postal_town") {
				address.Town = addressElem.Long_name
			} else if contains(addressElem.Types, "postal_code") {
				address.Postcode = addressElem.Long_name
			} else if contains(addressElem.Types, "country") {
				address.Country = addressElem.Long_name
			} else if contains(addressElem.Types, "administrative_area_level_2") {
				address.County = addressElem.Long_name
			}
		}
	} else {
		//TODO cope with error
		log.Fatal(err)
	}

	return address
}
