package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"GRPCS/Server_Stream/greetpb"
)

type server struct{

}

func (*server) Greet (ctx context.Context, req *greetpb.GreetRequest )(*greetpb.GreetResponse, error){
	fmt.Printf("Greet function was invoked with %v",req)
	firstname := req.GetGreeting().GetFirstName()
	result := "Hello" + firstname
	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res,nil

}

func (*server)GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error{
	fmt.Println("Streaming function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	for i:=0; i<10; i++{
		result := "Hello" + firstName + "number" + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 *time.Millisecond)
	}
	return nil

}





func main(){
	fmt.Println("Hello World")

    //lis,err := net.Listen(network string, dddress string)
	lis,err := net.Listen("tcp", "0.0.0.0:5051")
	if err!=nil {
		log.Fatalf("Failed to listen: %v",err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis) ; err!=nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}