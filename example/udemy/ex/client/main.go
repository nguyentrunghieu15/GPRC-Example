package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "example.com/example/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Sum(cc pb.CaculateSumClient, a int32, b int32) {
	res, err := cc.SumTowInteger(context.Background(), &pb.SumRequest{N1: a, N2: b})
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	fmt.Printf("Sum %v and %v: %v\n", a, b, res.Sum)
}

func SumArr(cc pb.CaculateSumClient, arrs []int32) {
	res, err := cc.SumArr(context.Background(), &pb.SumArrRequest{Arrs: arrs})
	if err != nil {
		log.Fatalln("Error: ", err)
	} else {
		fmt.Println("Sum of arrs: ", res.Sum)
	}
}

func Primes(cc pb.CaculateSumClient, n int32) {
	stream, err := cc.Primes(context.Background(), &pb.PrimesRequest{N: n})
	if err != nil {
		log.Fatalln("Error: ", err)
	} else {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln("Error: ", err)
			}
			fmt.Println(res.N)
		}
	}
}

func Avg(cc pb.CaculateSumClient, arrs []int32) {
	stream, err := cc.Avg(context.Background())
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	for _, i := range arrs {
		stream.Send(&pb.PrimesRequest{N: int32(i)})
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	fmt.Println("Avg ", arrs, " ", res.N)
}

func Max(cc pb.CaculateSumClient, arrs []int32) {
	stream, err := cc.Max(context.Background())
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	go func() {
		for _, i := range arrs {
			err := stream.Send(&pb.PrimesRequest{N: i})
			if err != nil {
				log.Fatalln("Error when sendding message: ", err)
			}
		}
		stream.CloseSend()
	}()

	c := make(chan (byte))

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				close(c)
				break
			}

			if err != nil {
				log.Fatalln("Error when sendding message: ", err)
			}

			fmt.Println(res.N)
		}
	}()

	<-c

}

func main() {
	addr := "localhost:5000"
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Can't connect to server with error", err)
	} else {
		log.Println("Connected to server")
	}
	defer conn.Close()
	cc := pb.NewCaculateSumClient(conn)
	Sum(cc, 1, 2)
	SumArr(cc, []int32{1, 2, 3, 4, 5})
	Primes(cc, 120)
	Avg(cc, []int32{1, 2, 3, 4, 5, 1})
	Max(cc, []int32{3, 14, 123, 12, 41, 23, 142})
}
