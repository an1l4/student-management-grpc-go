package serverHandler

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"sync"

	pb "github.com/an1l4/go-studentmgmt-grpc/studentmgmt"
)

var students []*pb.StudentList

type StudentManagementServer struct {
	pb.UnimplementedStudentManagementServer
	mu sync.Mutex
}

func InitStudents() {
	student1 := &pb.StudentList{Id: "1023", Name: "Alice", Rollnumber: 23, Class: 10, Section: "A", Hobby: []string{"sleeping", "eating"}}
	student2 := &pb.StudentList{Id: "0912", Name: "Bob", Rollnumber: 12, Class: 9, Section: "B", Hobby: []string{"reading", "playing"}}

	students = append(students, student1)
	students = append(students, student2)
}

func (s *StudentManagementServer) GetAllStudent(in *pb.Empty, stream pb.StudentManagement_GetAllStudentServer) error {
	log.Printf("Received: %v", in)
	for _, student := range students {
		if err := stream.Send(student); err != nil {
			return err
		}
	}
	return nil

}

func (s *StudentManagementServer) GetStudentBySection(ctx context.Context, in *pb.Section) (*pb.StudentList, error) {
	log.Printf("Received: %v", in)

	res := &pb.StudentList{}

	for _, student := range students {
		if student.GetSection() == in.GetValue() {
			res = student
			break
		} else {
			log.Printf("No student exist with this section")
		}
	}
	return res, nil
}

func (s *StudentManagementServer) GetStudentByClass(ctx context.Context, in *pb.Class) (*pb.StudentList, error) {
	log.Printf("Received: %v", in)

	res := &pb.StudentList{}

	for _, student := range students {
		if student.GetClass() == in.GetValue() {
			res = student
			break
		} else {
			log.Printf("No student exist in this class")
		}
	}
	return res, nil
}

func (s *StudentManagementServer) GetStudentByClassAndSection(in *pb.ClassSectionInfo, stream pb.StudentManagement_GetStudentByClassAndSectionServer) error {
	log.Printf("Received: %v", in.Classvalue)
	log.Printf("Received: %v", in.Sectionvalue)

	res := &pb.StudentList{}
	place := []*pb.StudentList{}

	for _, student := range students {
		if student.GetClass() == in.GetClassvalue() && student.GetSection() == in.GetSectionvalue() {
			res = student
			place = append(place, res)
			break

		} else {
			log.Printf("No student exist in this class and section")
		}
	}

	for _, i := range place {
		log.Printf("%v", i)
	}
	return nil
}

func (s *StudentManagementServer) CreateNewStudent(ctx context.Context, in *pb.StudentList) (*pb.Id, error) {
	log.Printf("Received:%v", in)
	//s.mu.Lock()
	//defer s.mu.Unlock()
	res := pb.Id{}

	for _, student := range students {
		if in.Name == student.Name {
			log.Fatalf("Student %v already exist", in.Name)
			break
		} else {
			res.Value = strconv.Itoa(rand.Intn(1000))
			in.Id = res.GetValue()

			students = append(students, in)

		}
	}
	return &res, nil

}

func (s *StudentManagementServer) UpdateStudent(ctx context.Context, in *pb.StudentList) (*pb.Status, error) {
	log.Printf("Received: %v", in)
	s.mu.Lock()
	defer s.mu.Unlock()

	res := pb.Status{}

	for index, student := range students {
		if student.GetId() == in.GetId() {
			students = append(students[:index], students[index+1:]...)
			in.Id = student.GetId()
			students = append(students, in)
			res.Value = 1
			break
		}
	}
	return &res, nil

}

func (s *StudentManagementServer) DeleteStudent(ctx context.Context, in *pb.Id) (*pb.Status, error) {
	log.Printf("Received: %v", in)

	res := pb.Status{}
	s.mu.Lock()
	defer s.mu.Unlock()

	for index, student := range students {
		if student.GetId() == in.GetValue() {
			students = append(students[:index], students[index+1:]...)
			res.Value = 1
			break
		} else {
			log.Printf("No student details in this id")
		}
	}

	return &res, nil
}
