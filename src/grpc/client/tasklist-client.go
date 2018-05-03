package main

import (
	"context"
	"flag"
	"fmt"
	"grpc/api"

	"google.golang.org/grpc"
)

var (
	address = flag.String("address", "localhost:10000", "The server address")
)

func main() {
	flag.Parse()

	fmt.Println(*address)

	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	taskServiceClient := taskservice.NewTaskServiceClient(conn)

	ctx := context.Background()

	task := taskservice.Task{
		Id:          0,
		Description: "Feed the cat",
		Owner: &taskservice.Owner{
			Id:    0,
			Name:  "Andrew Malcolm",
			Email: "arm@hat.com",
		},
	}

	_, err = taskServiceClient.AddTask(ctx, &task)
	if err != nil {
		panic(err)
	}

	void := taskservice.Void{}
	taskList, err := taskServiceClient.GetTaskList(ctx, &void)
	if err != nil {
		panic(err)
	}

	for _, task := range taskList.Task {
		fmt.Printf("Task: %v\n", task)
	}
}
