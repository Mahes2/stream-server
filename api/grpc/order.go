package grpc

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/codespade/stream-server/entity"
	"github.com/codespade/stream-server/pb"
)

var orderIds map[int64]bool

func (s *Server) ListOrder(stream pb.Order_ListOrderServer) error {
	orderIds = make(map[int64]bool)

	ticker := time.NewTicker(10 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				allOrders, newOrders := s.getOrders()
				if len(newOrders) > 0 {
					stream.Send(constructResponse(allOrders))
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	for {
		request, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			fmt.Println("err: ", err)
			fmt.Println(request)
			return err
		}

		fmt.Println(request)

		allOrders, _ := s.getOrders()
		stream.Send(constructResponse(allOrders))
	}
}

func (s *Server) getOrders() (allOrders []entity.Order, newOrders []entity.Order) {
	orders, err := s.OrderService.ListOrder(context.Background(), 0)
	if err != nil {
		fmt.Println("err: ", err)
	}
	for _, v := range orders {
		if !orderIds[v.Id] {
			newOrders = append(newOrders, v)
			fmt.Println("new order: ", v.Id)
			orderIds[v.Id] = true
		}
	}

	return orders, newOrders
}

func constructResponse(orders []entity.Order) *pb.ListOrderResponse {
	ordersPb := make([]*pb.ListOrderResponse_Order, 0)
	for _, v := range orders {
		ordersPb = append(ordersPb, &pb.ListOrderResponse_Order{
			Id:              v.Id,
			Quantity:        int32(v.Quantity),
			CustomerAddress: v.CustomerAddress,
		})
	}
	return &pb.ListOrderResponse{
		Order: ordersPb,
	}
}
