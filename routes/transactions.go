package routes

import (
  "github.com/gin-gonic/gin"
  "github.com/jmp3833/finance-backend/db"
  "github.com/jmp3833/finance-backend/models"
  "time"
  "fmt"
  "net/http"
)

func AddTransactionRoutes(r *gin.Engine) *gin.Engine {
  r.GET("/transactions/:bankName", func(c *gin.Context) {
    dbinst := db.GetDbInstance()
    bank := ParseBankName(c.Param("bankName"))
    afterDate := c.Query("afterDate")
    var transactions []models.Transaction
    var err, parseError error
    fmt.Print(afterDate)
    if (afterDate != "") {
      var parsedDate time.Time
      parsedDate, parseError = time.Parse(afterDate, models.DbTimeFormat)
      fmt.Print(parsedDate)
      transactions, err = db.GetTransactionsAfterDate(*dbinst, bank, parsedDate)
    } else {
      transactions, err = db.GetAllTransactions(*dbinst, bank)
    }
    if (err != nil || bank == models.NotImplemented || parseError != nil) {
      c.JSON(http.StatusNotFound, gin.H{
        "error": err,
        "description" : "Error fetching records for bankName=" + bank })
      }
      c.JSON(http.StatusOK, gin.H{
        "result": transactions,
        "count":  len(transactions),
      })
      defer dbinst.Close()
    })
    return r
  }

  func ParseBankName(bankName string) string {
    var bank string
    switch bankName {
    case "chase":
      bank = models.Chase
    case "bofa":
      bank = models.BankOfAmerica
    case "simple":
      bank = models.Simple
    default:
      bank = models.NotImplemented
    }
    return bank
  }
