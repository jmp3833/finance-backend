package routes

import (
  "github.com/gin-gonic/gin"
)
func AddHealthcheckRoutes(r *gin.Engine) *gin.Engine {
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })
  return r
}
