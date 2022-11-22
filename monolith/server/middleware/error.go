package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"

	"github.com/Drozd0f/ttto-go/monolith/service"
)

var handleErrors map[error]int = map[error]int{
	service.ErrUserAlreadyExists: http.StatusBadRequest,
	service.ErrUserInvalidData:   http.StatusBadRequest,
	service.ErrUserNotExists:     http.StatusNotFound,
	service.ErrGameInvalidId:     http.StatusBadRequest,
	service.ErrGameNotFound:      http.StatusNotFound,
	service.ErrGameInvalidState:  http.StatusBadRequest,
	service.ErrGameUserExists:    http.StatusConflict,
	service.ErrGameUserNotExists: http.StatusForbidden,
	service.ErrGameNotTurnUser:   http.StatusForbidden,
	service.ErrGameCellOccupied:  http.StatusForbidden,
}

func ErrorHandler(c *gin.Context) {
	c.Next()

	var validErr validation.Errors
	for _, err := range c.Errors {
		switch {
		case errors.As(err, &validErr):
			c.JSON(http.StatusBadRequest, validErr)
			return
		default:
			status, ok := handleErrors[err]
			if ok {
				c.JSON(status, gin.H{
					"message": err.Error(),
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": http.StatusText(http.StatusInternalServerError),
			})
		}
	}
}
