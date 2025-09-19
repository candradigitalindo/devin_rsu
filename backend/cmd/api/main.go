package main

import (
	"log"
	"os"

	"github.com/candra/saas-rs-backend/internal/config"
	"github.com/candra/saas-rs-backend/internal/db"
	"github.com/candra/saas-rs-backend/internal/server"
)

func main() {
	cfg := config.Load()
	pool, err := db.NewPool(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	app := server.New(cfg, pool)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}
