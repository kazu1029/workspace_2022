package graph

import "graphql/sample2/article/client"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	ArticleClient *client.Client
}
