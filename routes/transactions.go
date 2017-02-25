package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmp3833/finance-backend/db"
  "net/http"
)

func AddTransactionRoutes(r *gin.Engine) *gin.Engine {
	r.GET("/transactions/:bankName", func(c *gin.Context) {
		dbinst := db.GetDbInstance()
		bankName := c.Param("bankName")
		transactions, err := db.GetAllTransactions(*dbinst, bankName)
    if (err != nil) {
      c.JSON(http.StatusNotFound, gin.H{
        "error": err,
        "description" : "Error fetching records for bankName=" + bankName })
    }
		c.JSON(http.StatusOK, gin.H{
			"result": transactions,
			"count":  len(transactions),
		})
    defer dbinst.Close()
	})
	return r;
}
