package jobs

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

func SeedChaseDB(datafile string) {
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

		record := models.GetEmpowerModel(models.Bank{
			Transtype:   chasetrans[0],
			Description: chasetrans[3],
			Amount:      math.Abs(transamount),
			Date:        chasetrans[1]})

		dbinstance := db.GetDBInstance()
		db.AddRecord(record, dbinstance)
	}
}
