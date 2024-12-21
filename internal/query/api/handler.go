package api

import (
	"errors"
	"net/http"

	"github.com/ValeryBMSTU/web-11/pkg/vars"

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

func (srv *Server) Handler(e echo.Context) error {
	name := e.QueryParam("name")
	if name != "" {

		if len([]rune(name)) > srv.maxSize {
			return e.String(http.StatusBadRequest, "name is too large")
		}

		err := srv.uc.SetHelloMessage(name)
		if err != nil {
			if !errors.Is(err, vars.ErrAlreadyExist) {
				return e.String(http.StatusInternalServerError, err.Error())
			}
		}

		return e.String(http.StatusOK, "Hello, "+name+"!")
	} else {
		msg, err := srv.uc.FetchHelloMessage()
		if err != nil {
			return e.String(http.StatusInternalServerError, err.Error())
		}
		return e.String(http.StatusOK, "Hello, "+msg+"!")
	}
}
