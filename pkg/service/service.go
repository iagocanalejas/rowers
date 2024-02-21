package service

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"rowers/internal/db"
	d "rowers/templates/views/dashboard"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Service struct {
	port int
	db   db.Repository
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	server := &Service{
		port: port,
		db:   db.New(),
	}

	return &http.Server{
		Addr:         fmt.Sprintf(":%d", server.port),
		Handler:      server.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}

func (s *Service) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/dist", "dist")

	e.GET("/", s.dashboard)
	e.GET("/health", s.healthHandler)

	e.GET("/users", s.GetUsers)
	e.POST("/users", s.CreateUser)
	e.GET("/users/:user_id", s.GetUserById)
	e.DELETE("/users/:user_id", s.DeleteUser)

	e.GET("/users/:user_id/weights", s.GetUserWeights)
	e.POST("/users/:user_id/weights", s.AddWeight)
	e.DELETE("/users/:user_id/weights/:weight_id", s.DeleteWeight)

	e.GET("/assistances", s.GetAssistances)
	e.POST("/assistances", s.CreateAssistance)
	e.DELETE("/assistances/:assistance_id", s.DeleteAssistance)

	e.GET("/users/:user_id/assistances", s.GetUserAssistance)
	e.POST("/users/:user_id/assistances", s.AddUserAssistance)
	e.GET("/users/:user_id/assistances/:assistance_id", s.GetUserAssistanceById)
	e.DELETE("/users/:user_id/assistances/:assistance_id", s.DeleteUserAssistance)

	return e
}

func (s *Service) dashboard(c echo.Context) error {
	return d.Dashboard().Render(c.Request().Context(), c.Response().Writer)
}

func (s *Service) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}
