package main

import (
	"fmt"
	"github.com/jmp3833/finance/jobs"
	"os"
)

func main() {
	args := os.Args[1:]
	switch args[0] {
	case "chase-seed":
		seedChaseDB(args[1])
	default:
		fmt.Println("Incorrect program args")
	}
}

func seedChaseDB(filename string) {
	jobs.SeedChaseDB(filename)
}
