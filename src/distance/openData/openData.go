package openData

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"math"
	"strconv"
)

func GetDistance(p1, p2 string) (int, error) {

	p1 = strings.Replace(p1, " ", "", -1)
	p2 = strings.Replace(p2, " ", "", -1)
	

	filename := "/home/tszpinda/dev/code/go/openData/Data/CSV/ex.csv"
	f, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	fmt.Println(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	reader := csv.NewReader(f)

	fmt.Println("start")
	count := 0
	var r1 []string
	var r2 []string
	p1found := false
	p2found := false
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return 0, err
		}
		count++
		p := strings.Replace(record[0], " ", "", -1)
		if strings.Contains(p, p1) {
			r1 = record
			p1found = true
		}
		if strings.Contains(p, p2) {
			r2 = record
			p2found = true
		}
		if p1found && p2found {
			break
		}
	}
	//record:
	//Postcode,Positional_quality_indicator,Eastings,Northings,Country_code,NHS_regional_HA_code,NHS_HA_code,Admin_county_code,Admin_district_code,Admin_ward_code
	fmt.Println("Found: ", count, p1, p2, r1, r2)
	
	//simple calculation
	x1,_ := strconv.ParseFloat(r1[2], 64)
	y1,_ := strconv.ParseFloat(r1[3], 64)
	x2,_ := strconv.ParseFloat(r2[2], 64)
	y2,_ := strconv.ParseFloat(r2[3], 64)


	xd := math.Abs(x2 - x1)
	yd := math.Abs(y2 - y1)
	d := math.Sqrt(xd * xd + yd * yd)
	return int(d), nil
	        
}
