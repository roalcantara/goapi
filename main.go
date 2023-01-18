// main.go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/roalcantara/api/controllers"
	"github.com/roalcantara/api/db"
	"github.com/roalcantara/api/initializers"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func init() {
	initializers.LoadEnv()
	db.ConnectDB()
	db.Migrate()
}

var (
	r *gin.Engine
)

func main() {
	r = gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/api/check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Welcome to Golang with Gorm and Postgres"})
	})
	controllers.AddTaskRoutes(r)
	r.Run()
}
