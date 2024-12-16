package api

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (srv *Server) Login(e echo.Context) error {
	username := e.FormValue("username")
	password := e.FormValue("password")

	// Throws unauthorized error
	if username != "admin" || password != "admin" {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		"admin",
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func (srv *Server) GetQuery(e echo.Context) error {
	// get query name
	name := e.QueryParam("name")
	msg, err := srv.uc.FetchQuery(name)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, msg)
}

func (srv *Server) PostQuery(e echo.Context) error {
	// post query name
	name := e.QueryParam("name")
	err := srv.uc.InsertQuery(name)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.NoContent(http.StatusOK)
}
