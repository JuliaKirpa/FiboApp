package api

import (
	pb "FiboApp/api/lib"
	"FiboApp/internal"
	"google.golang.org/grpc"
	"net"
)

func GRPCRun(port string, serve *grpc.Server) error {
	lis, _ := net.Listen("tcp", "localhost:"+port)

	server := &internal.GServer{}
	pb.RegisterFiboAppServer(serve, server)

	if err := serve.Serve(lis); err != nil {
		return err
	}
	return nil
}
