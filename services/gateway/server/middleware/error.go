package middleware

import (
	"log"
	"net/http"

	"github.com/Drozd0f/ttto-go/services/gateway/service"
	"github.com/gin-gonic/gin"
)

var handleErrors map[error]int = map[error]int{
	service.ErrValidation:        http.StatusBadRequest,
	service.ErrUserAlreadyExists: http.StatusBadRequest,
}

func ErrorHandler(c *gin.Context) {
	c.Next()
	for _, err := range c.Errors {
		log.Println(err)
		status, ok := handleErrors[err]
		log.Println("status -", status, "ok", ok)
		if ok {
			c.JSON(status, gin.H{
				"message": err.Error(),
			})
			return
		}

		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusText(http.StatusInternalServerError),
		})
	}
}
