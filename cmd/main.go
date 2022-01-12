package main

import (
	"FiboApp/api"
	"FiboApp/internal"
	"log"
	"os"
)

func main() {
	srv := new(api.Server)

	if err := srv.Run(os.Getenv("PORT"), internal.GetInterval()); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}
}
