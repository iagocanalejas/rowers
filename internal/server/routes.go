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

	e.GET("/", s.dashboard)
	e.GET("/health", s.healthHandler)

	e.GET("/users", s.GetUsers)
	e.POST("/users", s.CreateUser)
	e.GET("/users/:user_id", s.GetUserByID)
	e.DELETE("/users/:user_id", s.DeleteUser)

	e.GET("/users/:user_id/weights", s.GetWeightsByUserId)
	e.POST("/users/:user_id/weights", s.AddWeight)
	e.DELETE("/users/:user_id/weights/:weight_id", s.DeleteWeight)

	e.GET("/assistances", s.GetAssistances)
	e.POST("/assistances", s.CreateAssistance)
	e.DELETE("/assistances/:assistance_id", s.DeleteAssistance)

	return e
}

func (s *Server) dashboard(c echo.Context) error {
	return views.Dashboard().Render(c.Request().Context(), c.Response().Writer)
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}
