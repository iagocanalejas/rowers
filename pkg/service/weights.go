package service

import (
	"errors"
	"net/http"

	"rowers/internal/db"
	u "rowers/templates/views/users"

	"github.com/labstack/echo/v4"
)

type UserAndWeightParams struct {
	UserID   int64 `query:"user_id" param:"user_id" json:"user_id"`
	WeightID int64 `query:"weight_id" param:"weight_id" json:"weight_id"`
}

// TODO: add date-range
func (s *Service) GetUserWeights(c echo.Context) error {
	params := new(UserParams)
	if err := c.Bind(params); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return s.getWeightsByUserID(c, params.UserID)
}

func (s *Service) CreateWeight(c echo.Context) error {
	body := new(db.Weight)
	if err := c.Bind(body); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if body.Weight < 30 || body.Weight > 200 {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("invalid weight"))
	}

	c.Logger().Info("adding weight to user %d", body.UserID)
	if _, err := s.db.CreateWeight(body.UserID, body.Weight); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return s.getWeightsByUserID(c, body.UserID)
}

func (s *Service) DeleteWeight(c echo.Context) error {
	params := new(UserAndWeightParams)
	if err := c.Bind(params); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Logger().Info("removing weight %d from user %d", params.WeightID, params.UserID)
	if err := s.db.DeleteWeight(params.UserID, params.WeightID); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return s.getWeightsByUserID(c, params.UserID)
}

func (s *Service) getWeightsByUserID(c echo.Context, userID int64) error {
	weights, err := s.db.GetWeightsByUserID(userID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return u.UserWeights(userID, weights).Render(c.Request().Context(), c.Response().Writer)
}
