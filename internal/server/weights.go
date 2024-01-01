package server

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"rowers/internal/database"
	"rowers/internal/views"

	"github.com/labstack/echo/v4"
)

type WeightForm struct {
	UserId string `json:"user_id"`
	Weight string `json:"weight"`
}

func validateForm(c echo.Context) (*database.UserWeight, error) {
	weight_data := new(WeightForm)
	if err := c.Bind(weight_data); err != nil {
		log.Println(err)
		return nil, err
	}

	userId, err := strconv.ParseInt(weight_data.UserId, 10, 64)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	weight, err := strconv.ParseFloat(weight_data.Weight, 64)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if weight < 30 || weight > 200 {
		return nil, errors.New("invalid weight")
	}

	return &database.UserWeight{UserId: userId, Weight: weight}, nil
}

func (s *Server) addWeight(c echo.Context) error {
	weight_data, err := validateForm(c)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.Println("adding weight to user")
	if err := s.db.AddWeight(weight_data.UserId, weight_data.Weight); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := s.db.GetUserById(weight_data.UserId)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if c.Request().Header.Get("Accept") == "application/json" {
		return c.JSON(http.StatusCreated, user)
	}
	return views.UserRow(*user).Render(c.Request().Context(), c.Response().Writer)
}
