package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Drozd0f/ttto-go/server/middleware"
	"github.com/Drozd0f/ttto-go/service"
	"github.com/gin-gonic/gin"
)

func (s *Server) registerGameHandlers(g *gin.RouterGroup) {
	gameG := g.Group("/games")
	{
		gameG.POST("", middleware.AuthRequired(), s.createGame)
		gameG.GET("", s.getGames)
		gameG.GET("/:gameID", s.getGame)
		gameG.PATCH("/:gameID/login", middleware.AuthRequired(), s.loginGame)
	}
}

func (s *Server) createGame(c *gin.Context) {
	g, err := s.service.CreateGame(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusOK, g)
}

func (s *Server) getGames(c *gin.Context) {
	games, err := s.service.GetGames(c.Request.Context(), c.Request.URL.Query())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusOK, games)
}

func (s *Server) getGame(c *gin.Context) {
	g, err := s.service.GetGameByID(c.Request.Context(), c.Param("gameID"))
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidId):
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		case errors.Is(err, service.ErrGameNotFound):
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": http.StatusText(http.StatusInternalServerError),
			})
		}
		return
	}

	c.JSON(http.StatusOK, g)
}

func (s *Server) loginGame(c *gin.Context) {
	fmt.Println(c.Param("gameID"))
	g, err := s.service.LoginGame(c.Request.Context(), c.Param("gameID"))
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidId):
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		case errors.Is(err, service.ErrGameNotFound):
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
		case errors.Is(err, service.ErrInvalidState):
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
		case errors.Is(err, service.ErrUserInGame):
			c.JSON(http.StatusConflict, gin.H{
				"message": err.Error(),
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": http.StatusText(http.StatusInternalServerError),
			})
		}
		return
	}

	c.JSON(http.StatusOK, g)
}
