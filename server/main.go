package main

import (
	"fmt"
	"log"
	"net"

	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

var posts = []Post{
	{
		Id:     1,
		Title:  "My awesome article 1",
		Author: "Quentin Neyrat",
	},
	{
		Id:     2,
		Title:  "My awesome article 2",
		Author: "Quentin Neyrat",
	},
	{
		Id:     3,
		Title:  "My awesome article 3",
		Author: "Quentin Neyrat",
	},
}

func (s *Server) ListPosts(empty *google_protobuf.Empty, stream PostService_ListPostsServer) error {
	for _, post := range posts {
		fmt.Printf("Send post #%d \n", post.GetId())
		if err := stream.Send(&post); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:4000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	g := grpc.NewServer()
	RegisterPostServiceServer(g, NewServer())
	g.Serve(lis)
}
