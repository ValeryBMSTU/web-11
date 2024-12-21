package api

import (
	"fmt"

	"github.com/ValeryBMSTU/web-11/internal/auth/provider"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	minPassword int
	maxPassword int
	minUsername int
	maxUsername int

	server  *echo.Echo
	r       *echo.Group
	address string

	uc Usecase
}

func NewServer(ip string, port int, minPassword, maxPassword, minUsername, maxUsername int, secret string, uc Usecase) *Server {
	api := Server{
		minPassword: minPassword,
		maxPassword: maxPassword,
		minUsername: minUsername,
		maxUsername: maxUsername,

		uc: uc,
	}

	api.server = echo.New()
	api.server.Use(middleware.Logger())
	api.server.Use(middleware.Recover())

	api.server.POST("/register", api.Register)
	api.server.POST("/login", api.Login)
	api.r = api.server.Group("/restricted")

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(provider.JWTClaims)
		},
		SigningKey: []byte(secret),
	}

	api.r.Use(echojwt.WithConfig(config))
	// api.r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey:  []byte("your-secret-key"),
	// 	TokenLookup: "header:Authorization",
	// 	AuthScheme:  "Bearer",
	// }))
	api.r.GET("", api.Restricted)

	api.address = fmt.Sprintf("%s:%d", ip, port)

	return &api
}

func (api *Server) Run() {
	api.server.Logger.Fatal(api.server.Start(api.address))
}
