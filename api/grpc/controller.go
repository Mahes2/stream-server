package grpc

import (
	"context"
	"fmt"
	"io"

	"github.com/codespade/stream-server/pb"
)

func (s *Server) VerifyHash(stream pb.Stream_VerifyHashServer) error {
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			fmt.Println(err)
			return err
		}

		md5Hash := s.Service.HashToMD5(request.Id)

		if request.Hash == md5Hash {
			fmt.Println("Hash for ID ", request.Id, "is VALID")
			response := &pb.VerifyHashResponse{
				Id:     request.Id,
				Status: "success",
			}
			err = stream.Send(response)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Hash for ID ", request.Id, "is INVALID")
			response := &pb.VerifyHashResponse{
				Id:     request.Id,
				Status: "failed",
			}
			err = stream.Send(response)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func (s *Server) BlockId(ctx context.Context, request *pb.BlockIdRequest) (response *pb.BlockIdResponse, err error) {
	resp, err := s.Service.BlockID(ctx, request.Id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("ID: ", resp.Id, "|Status: ", resp.Status)
	response = &pb.BlockIdResponse{
		Id:     resp.Id,
		Status: resp.Status,
	}

	return response, nil
}
