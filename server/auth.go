package server

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"

	"github.com/Drozd0f/ttto-go/models"
	"github.com/Drozd0f/ttto-go/repository"
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
		var validErr validation.Errors
		switch {
		case errors.As(err, &validErr):
			c.JSON(http.StatusBadRequest, err)
		case errors.Is(err, repository.ErrUniqueConstraint):
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "user already exist",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}
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
		var validErr validation.Errors
		switch {
		case errors.As(err, &validErr):
			c.JSON(http.StatusBadRequest, err)
		case errors.Is(err, repository.ErrNoResult):
			c.JSON(http.StatusNotFound, gin.H{
				"message": "user not exist",
			})
		case errors.Is(err, models.ErrInvalidData):
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid data",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Auth-Token": t,
	})
}
