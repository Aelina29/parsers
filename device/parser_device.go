package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type SuportedDevice struct {
	Retail_Branding string
	Marketing_Name  string
	Device          string
	Model           string //Card brands found with this MID/OEMID
}

func main() {
	var Marketing_Name string = "Galaxy A01"
	var Device string = "a01q"

	csvFile, _ := os.Open("supported_devices.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.LazyQuotes = true //игнорить "
	var Retail_Branding string = "not found"
	var dev SuportedDevice
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		dev = SuportedDevice{
			Retail_Branding: line[0],
			Marketing_Name:  line[1],
			Device:          line[2],
			Model:           line[3],
		}
		if (dev.Marketing_Name == Marketing_Name) && (dev.Device == Device) {
			Retail_Branding = dev.Retail_Branding
		}
		if Retail_Branding == "" {
			Retail_Branding = "unknown"
		}
	}
	fmt.Println(Retail_Branding)
}
