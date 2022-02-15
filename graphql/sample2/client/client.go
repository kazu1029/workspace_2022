package client

import (
	"graphql/sample2/article/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn *grpc.ClientConn
	Service pb.ArticleServiceClient
}

func NewClient(url string) (*Client, error) {
	// grpc.WithInsecure() is deprecated
	// see: https://pkg.go.dev/google.golang.org/grpc#WithInsecure
	// also: https://stackoverflow.com/questions/70482508/grpc-withinsecure-is-deprecated-use-insecure-newcredentials-instead
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	c := pb.NewArticleServiceClient(conn)
	return &Client{conn, c}, nil
}