package main

import (
	"FiboApp/api"
	pb "FiboApp/api/lib"
	"FiboApp/internal"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	go httpStart()
	grpcStart()
}

func httpStart() {
	HTTPsrv := new(api.Server)

	if err := HTTPsrv.Run(os.Getenv("PORT"), internal.GetInterval()); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}
}
func grpcStart() {
	lis, _ := net.Listen("tcp", "localhost:"+os.Getenv("GRPC_PORT"))
	server := &internal.GServer{}

	GRPCServ := grpc.NewServer()
	pb.RegisterFiboAppServer(GRPCServ, server)

	if err := GRPCServ.Serve(lis); err != nil {
		panic(err)
	}
}
