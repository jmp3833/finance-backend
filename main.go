package main

import (
	"bytes"
	"encoding/csv"
	"github.com/jmp3833/finance/db"
	"github.com/jmp3833/finance/models"
	"io"
	"io/ioutil"
	"log"
	"math"
	"strconv"
)

func main() {
	datafile := "/Users/jpeterson/workspace/go/src/github.com/jmp3833/finance/data/chase.csv"

	dat, err := ioutil.ReadFile(datafile)
	if err != nil {
		log.Fatal(err)
	}

	csvreader := csv.NewReader(bytes.NewReader(dat))

	//burn header line
	csvreader.Read()

	for {
		chasetrans, err := csvreader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		transamount, err := strconv.ParseFloat(chasetrans[4], 2)

		record := models.Chase{
			Transtype:   chasetrans[0],
			Description: chasetrans[3],
			Amount:      math.Abs(transamount),
			Date:        chasetrans[1]}

		dbinstance := db.GetDBInstance()
		db.AddRecord(record, dbinstance)
	}
}
