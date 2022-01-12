package main

import (
	"FiboApp/api"
	"FiboApp/internal"
	"github.com/bradfitz/gomemcache/memcache"
	"log"
	"os"
)

func main() {
	srv := new(api.Server)
	mc := memcache.New("memcached:" + os.Getenv("MC_PORT"))

	if err := srv.Run(os.Getenv("PORT"), internal.GetInterval(mc)); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}
}
