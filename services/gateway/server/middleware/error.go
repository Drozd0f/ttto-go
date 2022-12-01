package middleware

import (
	"errors"
	"net/http"

	"github.com/Drozd0f/ttto-go/services/gateway/service"
	"github.com/gin-gonic/gin"
)

var handleErrors = map[error]int{
	service.ErrValidation:        http.StatusBadRequest,
	service.ErrUserNotExists:     http.StatusNotFound,
	service.ErrUserAlreadyExists: http.StatusConflict,
}

func ErrorHandler(c *gin.Context) {
	c.Next()
	for _, err := range c.Errors {
		var e *gin.Error
		if errors.As(err, &e) {
			status, ok := handleErrors[e.Err]
			if ok {
				c.JSON(status, gin.H{
					"message": err.Error(),
				})
				return
			}
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusText(http.StatusInternalServerError),
		})
	}
}
