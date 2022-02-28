package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/kazu1029/workspace_2022/grpc/grpc-docker-todo/api/todopb"
	objectid "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

type server struct{}

var collection *mongo.Collection

type todoItem struct {
	ID       objectid.ObjectID `bson:"_id,omitempty"`
	AuthorID string            `bson:"author_id"`
	Content  string            `bson:""content`
	Title    string            `bson:"title"`
}

func main() {
	fmt.Println("===================== Todo API Start ======================")

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatalf("Failed to create mongo db client: err=%v", err)
	}
	if err = client.Connect(context.Background()); err != nil {
		log.Fatalf("Failed to connect mongo db: err=%v", err)
	}

	collection = client.Database("mydb").Collection("blog")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: err=%v", err)
	}

	s := grpc.NewServer()
	todopb.RegisterTodoServiceServer(s, &server{})

	go func() {
		fmt.Println("Starting Server...")
		err = s.Serve(lis)
		if err != nil {
			log.Fatalf("Failed to start Server: err=%v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Closing the listener")
	lis.Close()
	fmt.Println("===================== Todo API End ======================")
}

func (*server) CreateTodo(ctx context.Context, req *todopb.CreateTodoRequest) (*todopb.CreateTodoResponse, error) {
	fmt.Printf("Create Todo Item with %v\n", req)

	todo := req.GetTodo()
	data := todoItem{
		AuthorID: todo.GetAuthorId(),
		Title:    todo.GetTitle(),
		Content:  todo.GetContent(),
	}
	res, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return nil, err
	}

	oid, ok := res.InsertedID.(objectid.ObjectID)
	if !ok {
		return nil, fmt.Errorf("Cannot convert to OID")
	}
	data.ID = oid

	return &todopb.CreateTodoResponse{
		Todo: &todopb.Todo{
			Id:       data.ID.Hex(),
			AuthorId: data.AuthorID,
			Title:    data.Title,
			Content:  data.Content,
		},
	}, nil
}
