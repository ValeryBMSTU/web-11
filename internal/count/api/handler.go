package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func validateToken(token string) (bool, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8885/restricted", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		return false, err
	}

	return true, nil
}

func (srv *Server) jwtAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header.Get("Authorization") != "" {
			token := c.Request().Header.Get("Authorization")[7:]
			valid, err := validateToken(token)
			if err != nil || !valid {
				return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
			}
			return next(c)
		} else {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Вы не указали токен"})
		}
	}
}

func (srv *Server) GetCounter(e echo.Context) error {
	value, err := srv.uc.FetchCount()
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.String(http.StatusOK, strconv.Itoa(value))
}

func (srv *Server) PostCounter(e echo.Context) error {
	a, err := strconv.Atoi(e.FormValue("count"))
	if err != nil {
		e.Logger().Error(err)
		return e.String(http.StatusBadRequest, "это не число")
	}

	if a > srv.maxSize {
		return e.String(http.StatusBadRequest, "число слишком большое")
	}

	err = srv.uc.IncreaseCount(a)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.String(http.StatusOK, "OK!")
}
