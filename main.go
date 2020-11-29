package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{
		FieldMap: log.FieldMap{
			log.FieldKeyTime:  "@timestamp",
			log.FieldKeyMsg:   "message",
			log.FieldKeyLevel: "log.level",
			log.FieldKeyFile:  "log.origin.file.name",
			log.FieldKeyFunc:  "log.origin.function",
		}})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)

	log.SetReportCaller(true)
}

func main() {
	engine := gin.New()

	engine.GET("/healthz", func(c *gin.Context) {
		log.Info("calling healthz")
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	engine.GET("/ping", func(c *gin.Context) {
		log.Info("calling ping")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	engine.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")

		if name == "" {
			name = "world"
		}
		
		message := fmt.Sprintf("Hello %s", name)

		log.WithFields(log.Fields{
			"fields.tag": "hello",
			"user.id":    123,
			"user.name":  name,
			"tracing.trace.id": "aabc:222:33a",
			"tracing.transaction.id": "abc-2ww-abcd",
		}).Info(message)

		c.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	})

	engine.Run(":8080")
}
