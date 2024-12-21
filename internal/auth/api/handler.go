package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ValeryBMSTU/web-11/internal/auth/provider"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (srv *Server) Register(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}

	if len(username) < srv.minUsername || len(username) > srv.maxUsername {
		return echo.NewHTTPError(http.StatusUnauthorized, "Username should be "+strconv.Itoa(srv.minUsername)+"-"+strconv.Itoa(srv.maxUsername)+" length")
	}

	if len(password) < srv.minPassword || len(password) > srv.maxPassword {
		return echo.NewHTTPError(http.StatusUnauthorized, "Password should be "+strconv.Itoa(srv.minPassword)+"-"+strconv.Itoa(srv.maxPassword)+" length")
	}

	err := srv.uc.Register(username, password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Couldn't create account")
	}

	return c.JSON(http.StatusOK, "OK!")
}

func (srv *Server) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}

	token, err := srv.uc.Authenticate(username, password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error generating token")
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}

func (srv *Server) Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if user == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Token is missing or invalid"})
	}
	claims := user.Claims.(*provider.JWTClaims)
	log.Printf("Claims: %v", claims)
	username := claims.Username
	log.Printf("Username: %v", username)
	return c.JSON(http.StatusOK, map[string]string{"message": "Welcome " + username})
}
