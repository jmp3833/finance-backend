package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	chaseData, err := ioutil.ReadFile("../../../../data/chase.csv")
	if err != nil {
		fmt.Print(err)
	}

	chaseCsvReader := csv.NewReader(bytes.NewReader(chaseData))
	records, err := chaseCsvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(records)
}
