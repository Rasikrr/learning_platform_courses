package grpc

import (
	coursesS "github.com/Rasikrr/learning_platform_courses/internal/services/courses"
	pb "github.com/Rasikrr/learning_platform_courses/pkg/api/proto/courses"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCoursesServer
	coursesService coursesS.Service
}

func NewServer(
	srv *grpc.Server,
	coursesService coursesS.Service,
) {
	server := &server{
		coursesService: coursesService,
	}
	pb.RegisterCoursesServer(srv, server)
}
