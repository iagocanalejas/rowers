package server

import (
	"net/http"

	"rowers/internal/views"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/dist", "dist")

	e.GET("/", s.index)
	e.GET("/health", s.healthHandler)

	e.GET("/users", s.getUsers)
	e.POST("/users", s.createUser)

	e.GET("/users/:user_id", s.getUserById)
	e.DELETE("/users/:user_id", s.deleteUser)

	e.GET("/users/:user_id/weights", s.getWeights)
	e.POST("/users/:user_id/weights", s.addWeight)
	e.DELETE("/users/:user_id/weights/:weight_id", s.deleteWeight)

	return e
}

func (s *Server) index(c echo.Context) error {
	return views.Index().Render(c.Request().Context(), c.Response().Writer)
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}
