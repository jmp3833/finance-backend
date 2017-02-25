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
  r.GET("/transactions", func(c *gin.Context) {
    dbinstance := db.GetDbInstance()
    dbinstance.GetAllTransactions()
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r;
}
