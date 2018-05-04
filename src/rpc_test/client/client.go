package main

import (
	"flag"
	"fmt"
	"net/rpc"
	api "rpc_test/shared"
)

var (
	address = flag.String("address", "localhost:1234", "service address e.g. localhost:1234")
)

func main() {

	flag.Parse()

	client, err := rpc.Dial("tcp", *address)

	if err != nil {
		panic(err)
	}

	in := api.Task{0, "Feed the cat", api.Owner{}, 0}
	err = AddTask(client, &in, &api.Void{})

	if err != nil {
		panic(err)
	}

	tasks := []api.Task{}
	err = ListTasks(client, &api.Void{}, &tasks)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%d tasks in list\n", len(tasks))
	for _, task := range tasks {
		fmt.Printf("%v\n", task)
	}
}

// ListTasks : client helper func
func ListTasks(client *rpc.Client, in *api.Void, out *[]api.Task) error {
	return client.Call("TaskListServer.ListTasks", in, out)
}

// AddTask : client helper func
func AddTask(client *rpc.Client, in *api.Task, out *api.Void) error {
	return client.Call("TaskListServer.AddTask", in, out)
}
