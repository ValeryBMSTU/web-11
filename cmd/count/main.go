package main

import (
	"flag"
	"log"

	"github.com/ValeryBMSTU/web-11/internal/count/api"
	"github.com/ValeryBMSTU/web-11/internal/count/config"
	"github.com/ValeryBMSTU/web-11/internal/count/provider"
	"github.com/ValeryBMSTU/web-11/internal/count/usecase"
	_ "github.com/lib/pq"
)

func main() {
	configPath := flag.String("config-path", "./configs/count.yaml", "путь к файлу конфигурации")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	prv := provider.NewProvider(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBname)
	use := usecase.NewUsecase(cfg.Usecase.DefaultCount, prv)
	srv := api.NewServer(cfg.IP, cfg.Port, cfg.API.MaxCount, use)

	srv.Run()
}
