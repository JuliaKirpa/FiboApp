package internal

import (
	pb "FiboApp/api/lib"
	"FiboApp/pkg"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
)

func GetInterval() *gin.Engine {
	router := gin.Default()

	router.GET("/:from/:to", func(c *gin.Context) {
		from, err := strconv.Atoi(c.Param("from"))
		if err != nil {
			c.String(http.StatusBadRequest, "Use positive integer value for FROM\n")
		}
		to, err := strconv.Atoi(c.Param("to"))
		if err != nil {
			c.String(http.StatusBadRequest, "Use positive integer value for TO\n")
		}
		er := pkg.Validate(from, to)
		if er == "" {
			conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
			if err != nil {
				panic(err)
			}
			client := pb.NewFiboAppClient(conn)

			req := pb.Request{
				From: int32(from),
				To:   int32(to),
			}
			response, err1 := client.GRPCStart(context.Background(), &req)
			if err1 != nil {
				c.String(http.StatusBadRequest, "err from gRPC Start")
			}
			//JSONresp := pkg.Prepare(response, from)
			c.JSON(http.StatusOK, response)
		} else {
			c.String(http.StatusBadRequest, er)
		}
	})
	return router
}
