package server

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Drozd0f/ttto-go/service"
)

func (s *Server) GetUser(c *gin.Context) {
	u, err := s.service.GetUserByID(c.Request.Context(), c.Param("userId"))
	if err != nil {
		switch {
		case errors.Is(err, service.ErrUserNotFound):
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
		case errors.Is(err, service.ErrInvalidId):
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}
		return
	}
	c.JSON(http.StatusOK, u)
}
