package main

import (
	"fmt"
	"github.com/jmp3833/finance/jobs"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		printHelp()
		return
	}
	switch args[0] {
	case "chase-seed":
		handleSeedChaseInput()
	default:
		printHelp()
	}
}

func seedChaseDB(filename string) {
	jobs.SeedChaseDB(filename)
}

func printHelp() {
	fmt.Println("Accepted Command line args:")
	fmt.Println("chase-seed <path/to/csvfile> : Seed DB with CSV provided by chase bank export")
}

func handleSeedChaseInput() {
	args := os.Args[1:]
	if len(args) < 2 {
		printHelp()
		return
	}
	seedChaseDB(args[1])
}
