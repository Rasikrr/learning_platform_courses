package grpc

import (
	"context"
	pb "github.com/Rasikrr/learning_platform_courses/pkg/api/proto/courses"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) CreateCourse(ctx context.Context, req *pb.CreateCourseRequest) (*pb.EmptySuccessResponse, error) {
	if err := s.coursesService.CreateCourse(ctx, convertCreateCourseReq(req)); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating course: %v", err)
	}
	return &pb.EmptySuccessResponse{}, nil
}
