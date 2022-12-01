package middleware

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

func LoggingHandler(c *gin.Context) {
	c.Next()
	for _, err := range c.Errors {
		var e *gin.Error
		if errors.As(err, &e) {
			log.Println(e.Error())
			return
		}

		log.Println("Unexpected error:", err.Error())
	}
}
