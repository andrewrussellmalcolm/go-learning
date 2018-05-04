package main

import (
	"fmt"
	//"errors"

	"net"
	"net/rpc"
	api "rpc_test/shared"
)

type TaskListServer struct {
	taskList []api.Task
}

var taskListServer TaskListServer

func main() {

	taskListServer.taskList = append(taskListServer.taskList, api.Task{0, "Task 0", api.Owner{}, 0})
	server := rpc.NewServer()
	server.Register(&taskListServer)
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	server.Accept(listener)
}

func (tls *TaskListServer) AddTask(in *api.Task, out *api.Void) error {
	fmt.Println("AddTask")
	tls.taskList = append(tls.taskList, *in)
	return nil
}
func (tls *TaskListServer) ListTasks(in *api.Void, out *[]api.Task) error {
	fmt.Println("ListTasks")
	*out = tls.taskList
	return nil

}
func (tls *TaskListServer) DeleteTask(*api.Task, *api.Void) error {
	fmt.Println("DeleteTask")
	return nil
}
func (tls *TaskListServer) UpdateTask(*api.Task, *api.Void) error {
	fmt.Println("UpdateTask")
	return nil
}
