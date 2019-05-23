package main
import (
	"context"
	"fmt"
	"log"
	"io"
	"GRPCS/Server_Stream/greetpb"
	"google.golang.org/grpc"
)
func main(){
	fmt.Println("Hello I'm a client")
	conn,err :=grpc.Dial("localhost:5051", grpc.WithInsecure())
	if err != nil{
		log.Fatalf("could not connect: %v", err)
	}

	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	//fmt.Println("Created client: %f", c)
	doUnary(c)

	doServerStreaming(c)

}

func doUnary(c greetpb.GreetServiceClient){
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
		FirstName: "One",
		LastName: "Two",
	},
	}
	res,err :=c.Greet(context.Background(), req)
	if err!=nil {
		log.Fatalf("error while calling Greet RPC: %v", err)

	}
    log.Printf("Response from Greet: %v", res.Result)

}

func doServerStreaming(c greetpb.GreetServiceClient){
	fmt.Println("Starting to do a server streaming RPC")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Hay",
			LastName: "Man",
		},
	}

	resStream, err := 	c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyaTimes RPC: %v", err)
	}



	for {
		msg,err := resStream.Recv()

		if err ==io.EOF{
			break
		}

		if err != nil{
			log.Fatalf("error while reading stream %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v",msg.GetResult())
	}


}

