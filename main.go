package main

import (
	"fmt"
	"github.com/jmp3833/finance/jobs"
	"github.com/jmp3833/finance/models"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		printHelp()
		return
	}
	switch args[0] {
	case "simple-seed":
		handleSeed(models.SimpleTransaction{})
	case "chase-seed":
		handleSeed(models.ChaseTransaction{})
	case "fcu-seed":
		handleSeed(models.EmpowerFcuTransaction{})
	case "bofa-seed":
		handleSeed(models.BankOfAmericaTransaction{})
	default:
		printHelp()
	}
}

func printHelp() {
	fmt.Println("Accepted Command line args:")
	fmt.Println("chase-seed <path/to/csvfile>")
	fmt.Println("fcu-seed <path/to/csvfile>")
	fmt.Println("simple-seed <path/to/csvfile>")
	fmt.Println("bofa-seed<path/to/csvfile>")
}

func handleSeed(record models.Record) {
	args := os.Args[1:]
	if len(args) < 2 {
		printHelp()
		return
	}
	jobs.SeedDb(args[1], record)
}
