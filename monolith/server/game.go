package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"github.com/Drozd0f/ttto-go/monolith/models"
	"github.com/Drozd0f/ttto-go/monolith/server/middleware"
)

func (s *Server) registerGameHandlers(g *gin.RouterGroup) {
	gameG := g.Group("/games")
	{
		gameG.POST("", middleware.AuthRequired(), s.createGame)
		gameG.GET("", s.getGames)
		gameG.GET("/:gameID", s.getGame)
		gameG.PATCH("/:gameID", middleware.AuthRequired(), s.makeStep)
		gameG.GET("/:gameID/ws", s.ws)
		gameG.PATCH("/:gameID/login", middleware.AuthRequired(), s.loginGame)
	}
}

func (s *Server) createGame(c *gin.Context) {
	g, err := s.service.CreateGame(c.Request.Context())
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, g)
}

func (s *Server) getGame(c *gin.Context) {
	g, err := s.service.GetGameByID(c.Request.Context(), c.Param("gameID"))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, g)
}

func (s *Server) getGames(c *gin.Context) {
	games, err := s.service.GetGames(c.Request.Context(), c.Request.URL.Query())
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, games)
}

func (s *Server) makeStep(c *gin.Context) {
	var coord models.Coord

	if err := c.BindJSON(&coord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	g, err := s.service.MakeStep(c.Request.Context(), c.Param("gameID"), coord)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, g)
}

func (s *Server) loginGame(c *gin.Context) {
	g, err := s.service.LoginGame(c.Request.Context(), c.Param("gameID"))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, g)
}

func (s *Server) ws(c *gin.Context) {
	gameID := c.Param("gameID")
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	_, err := s.service.GetGameByID(ctx, gameID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	conn, err := s.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	defer conn.Close()

	go s.service.Subscribe(ctx, conn, gameID)

	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil || mt == websocket.CloseMessage {
			return
		}
		if ctx.Value("user") != nil {
			if err = s.service.HandleWSMessage(c.Request.Context(), gameID, msg); err != nil {
				_ = conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
				return
			}
		}

	}
}
