address-info
============

golang demo - how to google api and create own rest webservices:

1. returns full address from given postcode ie.: http://localhost:8080/ds/address/bs2
2. returns distance (in meters) between two given postcodes, ie.: http://localhost:8080/ds/distance/bs2/bs1


to get started (on unix):

git clone https://github.com/tszpinda/address-info.git
cd address-info
export GOPATH=`pwd`
go run src/main/addressInfo.go
