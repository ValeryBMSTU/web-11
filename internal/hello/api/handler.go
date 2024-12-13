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
		token := c.Request().Header.Get("Authorization")[7:]
		valid, err := validateToken(token)
		if err != nil || !valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
		}
		return next(c)
	}
}

// GetHello возвращает случайное приветствие пользователю
func (srv *Server) GetHello(e echo.Context) error {
	msg, err := srv.uc.FetchHelloMessage()
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, msg)
}

// PostHello Помещает новый вариант приветствия в БД
func (srv *Server) PostHello(e echo.Context) error {
	input := struct {
		Msg *string `json:"msg"`
	}{}

	err := e.Bind(&input)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	if input.Msg == nil {
		return e.String(http.StatusBadRequest, "msg is empty")
	}

	if len([]rune(*input.Msg)) > srv.maxSize {
		return e.String(http.StatusBadRequest, "hello message too large")
	}

	err = srv.uc.SetHelloMessage(*input.Msg)
	if err != nil {
		if errors.Is(err, vars.ErrAlreadyExist) {
			return e.String(http.StatusConflict, err.Error())
		}
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.String(http.StatusCreated, "OK")
}
