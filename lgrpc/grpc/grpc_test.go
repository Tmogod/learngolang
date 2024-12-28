package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	student_service "lgrpc/grpc/idl/my_proto"
	"testing"
)

func TestGrpc(t *testing.T) {
	conn, err := grpc.NewClient(
		"127.0.0.1:8082",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("连接出错:%v", err)
	}

	client := student_service.NewStudentServiceClient(conn)
	stu, err := client.GetStudentInfo(context.Background(), &student_service.Request{StudentId: "学生1"})
	if err != nil {
		t.Fatalf("RPC出错: %v", err)
	}
	fmt.Printf("name %s age %d height %f\n", stu.Name, stu.Age, stu.Height)
}
