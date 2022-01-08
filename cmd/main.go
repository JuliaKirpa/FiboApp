package main

import (
	"FiboApp/api"
	"FiboApp/internal"
	"log"
)

func main() {
	srv := new(api.Server)
	if err := srv.Run("8080", internal.GetInterval()); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}
}
