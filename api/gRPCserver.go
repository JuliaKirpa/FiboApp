package api

import (
	pb "FiboApp/api/lib"
	"google.golang.org/grpc"
	"net"
)

type gServer struct {
	pb.UnimplementedFiboAppServer
}

func GRPCRun(port string, serve *grpc.Server) error {
	lis, _ := net.Listen("tcp", "localhost:"+port)

	server := &gServer{}
	pb.RegisterFiboAppServer(serve, server)

	if err := serve.Serve(lis); err != nil {
		return err
	}
	return nil
}
