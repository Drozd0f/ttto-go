package server

import (
	"net/http"

	"github.com/Drozd0f/ttto-go/server/middleware"
	"github.com/gin-gonic/gin"
)

func (s *Server) registerGameHandlers(g *gin.RouterGroup) {
	gameG := g.Group("/games")
	{
		gameG.POST("", middleware.AuthRequired(), s.createGame)
		gameG.GET("", s.getGames)
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
