package main

import (
	"log"
	"net"

	"github.com/an1l4/go-studentmgmt-grpc/serverHandler"
	pb "github.com/an1l4/go-studentmgmt-grpc/studentmgmt"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	serverHandler.InitStudents()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterStudentManagementServer(s, &serverHandler.StudentManagementServer{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
