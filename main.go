package main

import (
	"fmt"
	"github.com/jmp3833/finance-backend/jobs"
	"github.com/jmp3833/finance-backend/models"
	"github.com/jmp3833/finance-backend/routes"
	"github.com/gin-gonic/gin"
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
	case "api":
	  startApi()
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
	fmt.Println("api start local api service")
}

func handleSeed(record models.Record) {
	args := os.Args[1:]
	if len(args) < 2 {
		printHelp()
		return
	}
	jobs.SeedDb(args[1], record)
  fmt.Print("All done!")
}

func startApi() {
  r := gin.Default()
  r = routes.AddHealthcheckRoutes(r)
  r = routes.AddTransactionRoutes(r)
  r.Run()
}
