package main

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"log"

	"github.com/kazu1029/workspace_2022/grpc/grpc-docker-todo/api/todopb"
	"google.golang.org/grpc"
)

func createTodoHandler(c todopb.TodoServiceClient) {
	fmt.Println("Creating the todo")

	authorId, err := makeRandomStr(20)
	if err != nil {
		log.Fatalf("Failed to generate authorId: err=%v", err)
	}
	title, err := makeRandomStr(30)
	if err != nil {
		log.Fatalf("Failed to generate title: err=%v", err)
	}
	content, err := makeRandomStr(50)
	if err != nil {
		log.Fatalf("Failed to generate content: err=%v", err)
	}

	todo := &todopb.Todo{
		AuthorId: authorId,
		Title:    title,
		Content:  content,
	}

	res, err := c.CreateTodo(context.Background(), &todopb.CreateTodoRequest{
		Todo: todo,
	})
	if err != nil {
		log.Fatalf("Error creating todo: %v", err)
	}
	fmt.Printf("res: %v\n", res)
}

func makeRandomStr(digit uint32) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 乱数を生成
	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error...")
	}

	// letters からランダムに取り出して文字列を生成
	var result string
	for _, v := range b {
		// index が letters の長さに収まるように調整
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}

func main() {
	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("server:50051", opts)
	if err != nil {
		fmt.Errorf("Error connecting to server: err=%v", err)
	}
	defer cc.Close()

	c := todopb.NewTodoServiceClient(cc)

	createTodoHandler(c)
}
