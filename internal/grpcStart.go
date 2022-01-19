package internal

import (
	pb "FiboApp/api/lib"
	"context"
	"math/big"
)

type GServer struct {
	pb.UnimplementedFiboAppServer
}

func (g *GServer) GRPCStart(ctx context.Context, inter *pb.Request) (*pb.Response, error) {
	var table []*big.Int
	table = mc.CheckMemcached(int(inter.From), int(inter.To))
	var jsonData []*pb.FiboVal

	for key, value := range table {
		numb := &pb.FiboVal{Id: int32(key) + inter.From, Value: value.Bytes()}
		jsonData = append(jsonData, numb)
	}
	resp := pb.Response{List: jsonData}

	return &resp, nil
}
