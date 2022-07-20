package grpc

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"

	pb "github.com/codespade/stream-server/pb"
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

		hasher := md5.New()
		hasher.Write([]byte(request.Id))
		data := hex.EncodeToString(hasher.Sum(nil))

		if request.Hash == string(data) {
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
	resp, err := s.Repository.BlockID(ctx, request.Id)
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