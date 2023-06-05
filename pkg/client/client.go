package client

import (
	"context"
	"fmt"

	pb "github.com/uyabpras/go-grpc-order-svc/pkg/proto/stubs"
	"google.golang.org/grpc"
)

type ProductServiceClient struct {
	Client pb.ProductServiceClient
}

func InitProductServiceClient(url string) ProductServiceClient {
	cc, err := grpc.Dial(url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("couldn't connect to:", err)
	}

	c := ProductServiceClient{
		Client: pb.NewProductServiceClient(cc),
	}

	return c
}

func (c *ProductServiceClient) FindOne(ProductID int64) (*pb.FindOneResponse, error) {
	req := &pb.FindOneRequest{
		Id: ProductID,
	}

	return c.Client.FindOne(context.Background(), req)
}

func (c *ProductServiceClient) DecreaseStock(ProductID int64, OrderID int64, Quantity int64) (*pb.DecreaseStockResponse, error) {
	req := &pb.DecreaseStockRequest{
		Id:       ProductID,
		OrderId:  OrderID,
		Quantity: Quantity,
	}

	return c.Client.DecreaseStock(context.Background(), req)
}
