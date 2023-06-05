package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/uyabpras/go-grpc-order-svc/pkg/client"
	"github.com/uyabpras/go-grpc-order-svc/pkg/db"
	"github.com/uyabpras/go-grpc-order-svc/pkg/models"
	"github.com/uyabpras/go-grpc-order-svc/pkg/proto/pb"
)

type Server struct {
	H          db.Handler
	Productsvc client.ProductServiceClient
}

func (s *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {

	product, err := s.Productsvc.FindOne(req.ProductID)

	fmt.Println(product.Data.Stock)

	if err != nil {
		return &pb.CreateOrderResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	} else if product.Status == http.StatusNotFound {
		return &pb.CreateOrderResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	} else if product.Data.Stock < req.Quantity {
		return &pb.CreateOrderResponse{
			Status: http.StatusConflict,
			Error:  "stock too low",
		}, nil
	}

	order := models.Order{
		Price:     product.Data.Price * req.Quantity,
		Quantity:  req.Quantity,
		ProductID: product.Data.Id,
		UserID:    req.UserID,
	}

	s.H.DB.Create(&order)

	res, err := s.Productsvc.DecreaseStock(req.ProductID, order.ID, order.Quantity)

	if err != nil {
		return &pb.CreateOrderResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	} else if res.Status == http.StatusConflict {
		s.H.DB.Delete(&models.Order{}, order.ID)
		return &pb.CreateOrderResponse{
			Status: http.StatusConflict,
			Error:  res.Error,
		}, nil
	}

	return &pb.CreateOrderResponse{
		Status: http.StatusCreated,
		ID:     order.ID,
	}, nil
}
