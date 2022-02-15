package main

import (
	"graphql/sample2/article/pb"
	"graphql/sample2/article/repository"
	"graphql/sample2/article/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	defer lis.Close()

	repository, err := repository.NewsqliteRepo()
	if err != nil {
		log.Fatalf("Failed to create sqlite repository: %v\n", err)
	}

	s := service.NewService(repository)

	server := grpc.NewServer()
	pb.RegisterArticleServiceServer(server, s)

	log.Println("Listening on port 50051...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
