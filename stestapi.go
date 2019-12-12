package main

import (
	pb "github.com/brotherlogic/stest/proto"
	"golang.org/x/net/context"
)

//Stest blah
func (s *Server) Stest(ctx context.Context, req *pb.Empty) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}
