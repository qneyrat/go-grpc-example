package main

import (
	"fmt"
	"io"
	"log"

	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

func printPosts(client PostServiceClient) {
	stream, err := client.ListPosts(context.Background(), &google_protobuf.Empty{})
	if err != nil {
		log.Fatalf("%v.ListPosts(_) = _, %v", client, err)
	}
	for {
		post, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListPosts(_) = _, %v", client, err)
		}

		fmt.Printf("Id: %d \n", post.GetId())
		fmt.Printf("Title: %s \n", post.GetTitle())
		fmt.Printf("Author: %s \n", post.GetAuthor())
	}
}

func main() {
	conn, err := grpc.Dial("localhost:4000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to start gRPC connection: %v", err)
	}
	defer conn.Close()

	client := NewPostServiceClient(conn)

	printPosts(client)
}
