package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	pb "github.com/varnit-ta/gRPC-Practice/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(":4000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloWorldServiceClient(conn)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter your name (or type 'exit' to quit): ")
		if !scanner.Scan() {
			break
		}
		name := strings.TrimSpace(scanner.Text())
		if strings.ToLower(name) == "exit" {
			fmt.Println("Exiting client.")
			break
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		r, err := c.SayHello(ctx, &pb.HelloWorldRequest{Name: name})
		cancel()
		if err != nil {
			log.Printf("could not greet: %v", err)
			continue
		}
		fmt.Printf("Greeting: %s\n", r.GetMessage())
	}
}
