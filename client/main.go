package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/brandon2255p/restaurant/restaurant-svc/proto"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "brandon"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("%s-%d", defaultName, i)
		r, err := c.Hello(ctx, &pb.Request{Name: name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.Msg)
	}

	log.Println("Exiting")
}
