package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gorilla/websocket"

	"github.com/Drozd0f/ttto-go/models"
	"github.com/Drozd0f/ttto-go/server/middleware"
	"github.com/Drozd0f/ttto-go/service"
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
		var validErr validation.Errors
		switch {
		case errors.As(err, &validErr):
			c.JSON(http.StatusBadRequest, err)
		case errors.Is(err, service.ErrInvalidId),
			errors.Is(err, service.ErrInvalidState):
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		case errors.Is(err, service.ErrGameNotFound):
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
		case errors.Is(err, service.ErrUserNotInGame),
			errors.Is(err, service.ErrNotYourTurn),
			errors.Is(err, models.ErrCellOccupied):
			c.JSON(http.StatusForbidden, gin.H{
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

func (s *Server) ws(c *gin.Context) {
	gameID := c.Param("gameID")
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	_, err := s.service.GetGameByID(ctx, gameID)
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
				conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
				return
			}
		}

	}
}
