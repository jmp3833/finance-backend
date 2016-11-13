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

func SeedEmpowerDb(datafile string) {
	dat, err := ioutil.ReadFile(datafile)
	if err != nil {
		log.Fatal(err)
	}

	csvreader := csv.NewReader(bytes.NewReader(dat))

	//burn header line
	csvreader.Read()

	for {
		fcuTrans, err := csvreader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		transamount, err := strconv.ParseFloat(fcuTrans[4], 2)

		record := models.GetEmpowerModel(models.Bank{
			Transtype:   fcuTrans[3],
			Description: fcuTrans[8],
			Amount:      math.Abs(transamount),
			Date:        fcuTrans[1]})

		dbinstance := db.GetDBInstance()
		db.AddRecord(record, dbinstance)
		dbinstance.Close()
	}
}
