package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/jmp3833/finance/models"
	"io"
	"io/ioutil"
	"log"
	"strconv"
)

func main() {
	dat, err := ioutil.ReadFile("./data/chase.csv")
	if err != nil {
		log.Fatal(err)
	}

	csvreader := csv.NewReader(bytes.NewReader(dat))

	for {
		chasetrans, err := csvreader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		transamount, err := strconv.ParseFloat(chasetrans[4], 32)

		record := models.Chase{
			Transtype:   chasetrans[0],
			Description: chasetrans[3],
			Amount:      transamount,
			Date:        chasetrans[1]}

		fmt.Printf("%+v\n", record)
	}
}
