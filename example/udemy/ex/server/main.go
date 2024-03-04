package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	pb "example.com/example/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.CaculateSumServer
}

func (Server) SumTowInteger(co context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Println("Recived request SumTowInteger", in)
	return &pb.SumResponse{Sum: in.N1 + in.N2}, nil
}

func (Server) SumArr(context context.Context, in *pb.SumArrRequest) (*pb.SumResponse, error) {
	log.Println("Recived request SumArr", in)
	var sum int32 = 0
	for _, i := range in.GetArrs() {
		sum += int32(i)
	}
	return &pb.SumResponse{Sum: sum}, nil
}

func (Server) Primes(in *pb.PrimesRequest, stream pb.CaculateSum_PrimesServer) error {
	log.Println("Recived request Primes: ", in.N)
	var temp = in.N
	var k int32 = 2
	for temp != 1 {
		if temp%k == 0 {
			stream.Send(&pb.PrimesResponse{N: k})
			temp /= k
			continue
		}
		k++
	}
	return nil
}

func (Server) Avg(stream pb.CaculateSum_AvgServer) error {
	log.Println("Recived request Avg")
	var sum float32 = 0
	n := 0
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			if n == 0 {
				log.Fatalln("Error: ", fmt.Errorf("Empty data"))
				return fmt.Errorf("Empty data")
			}
			stream.SendAndClose(&pb.AvgResponse{N: sum / float32(n)})
			break
		}
		if err != nil {
			log.Fatalln("Error: ", err)
			return err
		}
		sum += float32(res.N)
		n++
	}
	return nil
}

func main() {
	addr := "0.0.0.0:5000" // Address of server

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Cant not listen on address %v\n,Error:%v\n", addr, err)
	}

	s := grpc.NewServer()

	pb.RegisterCaculateSumServer(s, &Server{})

	log.Println("Creating gRPC Server on ", addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalln("Cant not create gRPC server", err)
	}
	defer lis.Close()
}
