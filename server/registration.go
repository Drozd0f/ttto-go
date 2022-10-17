package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Drozd0f/ttto-go/models"
)

func (s *Server) Reg(c *gin.Context) {
	// c.Request.Context()
	var u models.User

	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := s.service.Reg(&u); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusBadRequest, u)
}
