package clientHandler

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/an1l4/go-studentmgmt-grpc/studentmgmt"
)

func RunGetAllStudent(client pb.StudentManagementClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	req := &pb.Empty{}
	stream, err := client.GetAllStudent(ctx, req)
	if err != nil {
		log.Fatalf("%v.GeAllStudents(_) = _, %v", client, err)
	}

	for {
		row, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GeAllStudents(_) = _, %v", client, err)
		}
		log.Printf("StudentDetails: %v", row)
	}

}

func RunGetStudentBySection(client pb.StudentManagementClient, studentsection string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	req := &pb.Section{Value: studentsection}

	res, err := client.GetStudentBySection(ctx, req)

	if err != nil {
		log.Fatalf("%v.GetStudentBySection(_) = _, %v", client, err)
	}
	log.Printf("StudentDetails: %v", res)
}

func RunGetStudentByClass(client pb.StudentManagementClient, studentclass int32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	req := &pb.Class{Value: studentclass}
	res, err := client.GetStudentByClass(ctx, req)

	if err != nil {
		log.Fatalf("%v.GetStudentByClass(_) = _, %v", client, err)

	}
	log.Printf("StudentDetails: %v", res)
}

func RunGetStudentByClassAndSection(client pb.StudentManagementClient, classvalue int32, sectionvalue string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	req := &pb.ClassSectionInfo{Classvalue: classvalue, Sectionvalue: sectionvalue}

	res, err := client.GetStudentByClassAndSection(ctx, req)
	if err != nil {
		log.Fatalf("%v.GetStudentByClassAndSection(_) = _, %v", client, err)
	}
	log.Printf("Students :%v", res)
}

func RunCreateNewStudent(client pb.StudentManagementClient, name string, rollnumber int32, class int32, section string, hobby []string) {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)

	defer cancel()

	req := &pb.StudentList{Name: name, Rollnumber: rollnumber, Class: class, Section: section, Hobby: hobby}

	res, err := client.CreateNewStudent(ctx, req)

	if err != nil {
		log.Fatalf("unable to create new student %v", err)
	}

	if res.GetValue() != "" {
		log.Printf("New student created successfully and id is %v", res)
	} else {
		log.Printf("Creating new student is failed")
	}

}

func RunUpdateStudent(client pb.StudentManagementClient, studentid string, name string, rollnumber int32, class int32, section string, hobby []string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	req := &pb.StudentList{Id: studentid, Name: name, Rollnumber: rollnumber, Class: class, Section: section, Hobby: hobby}

	res, err := client.UpdateStudent(ctx, req)

	if err != nil {
		log.Fatalf("unable to update student details %v", err)
	}
	if int(res.GetValue()) == 1 {
		log.Printf("Student Details Updated Successfully")
	} else {
		log.Printf("Student Details updation failed")
	}
}

func RunDeleteStudent(client pb.StudentManagementClient, studentid string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	req := &pb.Id{Value: studentid}

	res, err := client.DeleteStudent(ctx, req)

	if err != nil {
		log.Fatalf("unable to delete student details  %v", err)
	}
	if int(res.GetValue()) == 1 {
		log.Printf("Student deleted successfully")
	} else {
		log.Printf("Student deletion failed")
	}
}
