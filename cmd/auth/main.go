package main

import (
	"flag"
	"log"

	"github.com/ValeryBMSTU/web-11/internal/auth/api"
	"github.com/ValeryBMSTU/web-11/internal/auth/config"
	"github.com/ValeryBMSTU/web-11/internal/auth/provider"
	"github.com/ValeryBMSTU/web-11/internal/auth/usecase"
	_ "github.com/lib/pq"
)

func main() {
	configPath := flag.String("config-path", "./configs/auth.yaml", "путь к файлу конфигурации")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	prv := provider.NewProvider(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBname)
	jwt_prv := provider.NewJWTProvider(cfg.JWT.Secret)
	use := usecase.NewUsecase(cfg.Usecase.DefaultMessage, prv, jwt_prv)
	srv := api.NewServer(cfg.IP, cfg.Port, cfg.API.MinPasswordSize, cfg.API.MaxPasswordSize, cfg.API.MinUsernameSize, cfg.API.MaxUsernameSize, cfg.JWT.Secret, use)

	srv.Run()
}
