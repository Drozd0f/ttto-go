package server

import (
	"net/http"

	"github.com/Drozd0f/ttto-go/gen/proto/auth"
	"github.com/gin-gonic/gin"
)

func (s *Server) registerAuthHandlers(g *gin.RouterGroup) {
	authG := g.Group("/auth")
	{
		authG.POST("/reg", s.registration)
		authG.POST("/log", s.login)
	}
}

func (s *Server) registration(c *gin.Context) {
	var u *auth.CreateUserRequest

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := s.service.CreateUser(c.Request.Context(), u); err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "OK",
	})
}

func (s *Server) login(c *gin.Context) {
	var u *auth.LoginUserRequest

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	token, err := s.service.LoginUser(c.Request.Context(), u)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
