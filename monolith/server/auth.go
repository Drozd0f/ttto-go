package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Drozd0f/ttto-go/monolith/models"
)

func (s *Server) registerAuthHandlers(g *gin.RouterGroup) {
	authG := g.Group("/auth")
	{
		authG.POST("/reg", s.reg)
		authG.POST("/login", s.login)
	}
}

func (s *Server) reg(c *gin.Context) {
	var u models.User

	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := s.service.Reg(c.Request.Context(), &u); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "OK",
	})
}

func (s *Server) login(c *gin.Context) {
	var u models.User

	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	t, err := s.service.Login(c.Request.Context(), &u)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Auth-Token": t,
	})
}
