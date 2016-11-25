package jobs

import (
	"bytes"
	"encoding/csv"
	"github.com/jmp3833/finance/db"
	"github.com/jmp3833/finance/models"
	"io"
	"io/ioutil"
	"log"
)

func SeedDb(datafile string, record models.Record) {
	f, err := ioutil.ReadFile(datafile)

	if err != nil {
		log.Fatal(err)
	}

	csvreader := csv.NewReader(bytes.NewReader(f))

	//burn header line
	csvreader.Read()
	dbinstance := db.GetDbInstance()
	defer dbinstance.Close()

	for {
		line, err := csvreader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		defer dbinstance.Close()

		db.InsertTransaction(dbinstance, record.ParseTransactionFromCsvLine(line))
	}
}