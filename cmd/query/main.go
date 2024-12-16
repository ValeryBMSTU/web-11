package main

import (
	"flag"
	"log"

	"github.com/ValeryBMSTU/web-11/internal/query/api"
	"github.com/ValeryBMSTU/web-11/internal/query/config"
	"github.com/ValeryBMSTU/web-11/internal/query/provider"
	"github.com/ValeryBMSTU/web-11/internal/query/usecase"
	_ "github.com/lib/pq"
)

func main() {
	// Считываем аргументы командной строки
	configPath := flag.String("config-path", "/home/user/Documents/lab11/web-11/configs/hello_example.yaml", "путь к файлу конфигурации")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	prv := provider.NewProvider(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBname)
	use := usecase.NewUsecase(cfg.Usecase.DefaultMessageQuery, prv)
	srv := api.NewServer(cfg.IP, cfg.Port, cfg.API.MaxMessageSize, use)

	srv.Run()
}
