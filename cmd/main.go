package main

import (
	"fmt"
	"log"
	"net"

	"github.com/uyabpras/go-grpc-order-svc/pkg/client"
	"github.com/uyabpras/go-grpc-order-svc/pkg/config"
	"github.com/uyabpras/go-grpc-order-svc/pkg/db"
	"github.com/uyabpras/go-grpc-order-svc/pkg/proto/pb"
	"github.com/uyabpras/go-grpc-order-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.Loadconfig()

	if err != nil {
		log.Fatalln("failed to load config: ", err)
	}

	h := db.Init(c.Dburl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("failed to listen: ", err)
	}

	productsrvc := client.InitProductServiceClient(c.ProductSvc)

	if err != nil {
		log.Fatalln("failed to listen: ", err)
	}

	fmt.Println("Order svc: ", c.Port)

	s := services.Server{
		H:          h,
		Productsvc: productsrvc,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
