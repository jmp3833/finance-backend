package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmp3833/finance-backend/db"
)

func AddRoutes(r *gin.Engine) *gin.Engine {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r;
	r.GET("/transactions/:bankName", func(c *gin.Context) {
		dbinst := db.GetDbInstance()
		bankName := c.Param("bankName")
		transactions := db.GetAllTransactions(dbinst, bankName)
		c.JSON(http.StatusOK, gin.H{
			"result": transactions,
			"count":  len(transactions),
		})
	})
	return r;
}
