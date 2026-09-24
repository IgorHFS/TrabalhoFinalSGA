package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"sga/internal/config"
	"sga/routes"
)

func main() {
	db, err := config.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	routes.Register(r, db, value("APP_VERSION", "1.0.0"))
	if err := r.Run(":" + value("SERVER_PORT", "8080")); err != nil {
		log.Fatal(err)
	}
}

func value(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
