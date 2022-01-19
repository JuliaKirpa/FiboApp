package main

import (
	"FiboApp/api"
	"FiboApp/internal"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {
	HTTPsrv := new(api.Server)

	if err := HTTPsrv.Run(os.Getenv("PORT"), internal.GetInterval()); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}

	GRPCServ := grpc.NewServer()

	if err := api.GRPCRun(os.Getenv("GRPC_PORT"), GRPCServ); err != nil {
		log.Fatalf("error while running gRPC server: %s", err.Error())
	}
}
