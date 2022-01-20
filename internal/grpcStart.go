package internal

import (
	pb "FiboApp/api/lib"
	"context"
	"fmt"
	"math/big"
)

type GServer struct {
	pb.UnimplementedFiboAppServer
}

func (g *GServer) GRPCStart(ctx context.Context, inter *pb.Request) (*pb.Response, error) {
	var table []*big.Int
	var jsonData []*pb.FiboVal

	table = mc.CheckMemcached(int(inter.GetFrom()), int(inter.GetTo()))

	for key, value := range table {
		numb := pb.FiboVal{Id: int32(key) + inter.GetFrom(), Value: value.Bytes()}
		fmt.Println(numb.Value)
		jsonData = append(jsonData, numb)
	}

	resp := pb.Response{List: jsonData}
	
	return &resp, nil
}
