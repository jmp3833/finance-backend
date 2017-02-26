package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmp3833/finance-backend/db"
	"github.com/jmp3833/finance-backend/models"
  "net/http"
)

func AddTransactionRoutes(r *gin.Engine) *gin.Engine {
	r.GET("/transactions/:bankName", func(c *gin.Context) {
		dbinst := db.GetDbInstance()
    var bank string
		switch c.Param("bankName") {
      case "chase":
        bank = models.Chase
      case "bofa":
        bank = models.BankOfAmerica
      case "simple":
        bank = models.Simple
      default:
        bank = models.NotImplemented
    }
		transactions, err := db.GetAllTransactions(*dbinst, bank)
    if (err != nil || bank == models.NotImplemented) {
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
	return r;
}
