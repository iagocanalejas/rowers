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
	Weight string `json:"weight"`
}

func (s *Server) getWeights(c echo.Context) error {
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	weights, err := s.db.GetWeights(user_id)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return views.UserWeights(user_id, weights).Render(c.Request().Context(), c.Response().Writer)
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
	log.Println(user.Weight)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if c.Request().Header.Get("Accept") == "application/json" {
		return c.JSON(http.StatusCreated, user)
	}
	return views.UserRow(*user).Render(c.Request().Context(), c.Response().Writer)
}

func (s *Server) deleteWeight(c echo.Context) error {
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	weight_id, err := strconv.ParseInt(c.Param("weight_id"), 10, 64)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.Printf("removing weight %d from user %d", weight_id, user_id)
	if err := s.db.DeleteUserWeight(user_id, weight_id); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

// TODO: search where to put this
func validateForm(c echo.Context) (*database.UserWeight, error) {
	weight_data := new(WeightForm)
	if err := c.Bind(weight_data); err != nil {
		log.Println(err)
		return nil, err
	}

	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
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
