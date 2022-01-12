package main

import (
	"FiboApp/api"
	"FiboApp/internal"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error load env variable: %s", err.Error())
	}

	srv := new(api.Server)

	if err := srv.Run(os.Getenv("PORT"), internal.GetInterval()); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}
}
