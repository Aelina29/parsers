package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type SD struct {
	Company string
	MID     string
	OEMID   string
	Brands  string //Card brands found with this MID/OEMID
}

func main() {
	var MID string = "0x00009c"
	var OEMID string = "BEs"

	csvFile, _ := os.Open("sd_cards.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.LazyQuotes = true //игнорить "
	var Company string = "not found"
	var card SD
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		card = SD{
			Company: line[0],
			MID:     line[1],
			OEMID:   line[2],
			Brands:  line[3],
		}
		if (card.MID == MID) && (card.OEMID == OEMID) {
			Company = card.Company
		}
		if Company == "" {
			Company = "unknown"
		}
	}
	fmt.Println(Company)
}
