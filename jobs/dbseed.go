package jobs

import (
	"bytes"
	"encoding/csv"
	"github.com/jmp3833/finance-backend/db"
	"github.com/jmp3833/finance-backend/models"
	"io"
  "fmt"
	"io/ioutil"
	"log"
)

func SeedDb(datafile string, record models.Record) {
	f, err := ioutil.ReadFile(datafile)
	if err != nil {
    fmt.Print("Unable to read CSV input file:")
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
		db.InsertTransaction(dbinstance, record.ParseTransactionFromCsvLine(line))
	}
}
