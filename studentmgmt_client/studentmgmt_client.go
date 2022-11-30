package main

import (
	"log"

	"github.com/an1l4/go-studentmgmt-grpc/clientHandler"
	pb "github.com/an1l4/go-studentmgmt-grpc/studentmgmt"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect %v", err)
	}
	defer conn.Close()
	client := pb.NewStudentManagementClient(conn)

	clientHandler.RunGetAllStudent(client)
	//clientHandler.RunGetStudentBySection(client,"A")
	//clientHandler.RunGetStudentByClass(client,9)
	//clientHandler.RunCreateNewStudent(client,"Anila",56,8,"C",[]string{"painting","dancing"})
	//clientHandler.RunGetStudentByClassAndSection(client,10,"A")
	//clientHandler.RunUpdateStudent(client,"81","Anna",23,10,"A",[]string{"sleeping","eating"})
	//clientHandler.RunDeleteStudent(client, "1023")
}
