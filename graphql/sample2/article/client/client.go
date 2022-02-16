package client

import (
	"context"
	"graphql/sample2/article/pb"
	"graphql/sample2/graph/model"
	"io"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	Conn    *grpc.ClientConn
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

func (c *Client) CreateArticle(ctx context.Context, input *pb.ArticleInput) (*model.Article, error) {
	res, err := c.Service.CreateArticle(
		ctx,
		&pb.CreateArticleRequest{ArticleInput: input},
	)
	if err != nil {
		return nil, err
	}

	return &model.Article{
		ID:      int64(res.Article.Id),
		Author:  res.Article.Author,
		Title:   res.Article.Title,
		Content: res.Article.Content,
	}, nil
}

func (c *Client) ReadArticle(ctx context.Context, id int64) (*model.Article, error) {
	res, err := c.Service.ReadArticle(ctx, &pb.ReadArticleRequest{Id: id})
	if err != nil {
		return nil, err
	}

	return &model.Article{
		ID:      int64(res.Article.Id),
		Author:  res.Article.Author,
		Title:   res.Article.Title,
		Content: res.Article.Content,
	}, nil
}

func (c *Client) UpdateArticle(ctx context.Context, id int64, input *pb.ArticleInput) (*model.Article, error) {
	res, err := c.Service.UpdateArticle(ctx, &pb.UpdateArticleRequest{Id: id, ArticleInput: input})
	if err != nil {
		return nil, err
	}

	return &model.Article{
		ID:      int64(res.Article.Id),
		Author:  res.Article.Author,
		Title:   res.Article.Title,
		Content: res.Article.Content,
	}, nil
}

func (c *Client) DeleteArticle(ctx context.Context, id int64) (int64, error) {
	res, err := c.Service.DeleteArticle(ctx, &pb.DeleteArticleRequest{Id: id})
	if err != nil {
		return 0, err
	}

	return res.Id, nil
}

func (c *Client) ListArticle(ctx context.Context) ([]*model.Article, error) {
	res, err := c.Service.ListArticle(ctx, &pb.ListArticleRequest{})
	if err != nil {
		return nil, err
	}
	var articles []*model.Article
	for {
		r, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		articles = append(articles, &model.Article{
			ID:      int64(r.Article.Id),
			Author:  r.Article.Author,
			Title:   r.Article.Title,
			Content: r.Article.Content,
		})
	}

	return articles, nil
}
