package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	student_service "lgrpc/grpc/idl/my_proto"
	"net"
)

type StudentServer struct {
	student_service.UnimplementedStudentServiceServer
}

func (s StudentServer) GetStudentInfo(ctx context.Context, in *student_service.Request) (*student_service.Student, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err:", err)
		}
	}()
	if len(in.StudentId) == 0 {
		return nil, errors.New("学生ID不能为空")
	}
	student := &student_service.Student{
		Name:   "张三",
		Age:    18,
		Height: 1.75,
	}

	return student, nil
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8082")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	student_service.RegisterStudentServiceServer(server, new(StudentServer))
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
