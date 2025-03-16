package grpc

import (
	"context"
	pb "github.com/Rasikrr/learning_platform_courses/pkg/api/proto/courses"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) GetCoursesByParams(ctx context.Context, req *pb.GetCoursesByParamsRequest) (*pb.GetCoursesByParamsResponse, error) {
	out, err := s.coursesService.GetCoursesByParams(ctx, convertGetCoursesByParamsReq(req))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting courses by params: %v", err)
	}
	resp := &pb.GetCoursesByParamsResponse{
		Courses: convertCoursesToPb(out),
	}
	return resp, nil
}

func (s *server) GetCourseByID(ctx context.Context, req *pb.GetCourseByIDRequest) (*pb.GetCourseByIDResponse, error) {
	out, err := s.coursesService.GetCourseByID(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting course by id: %v", err)
	}
	return &pb.GetCourseByIDResponse{Course: convertCourseToPb(out)}, nil
}

func (s *server) GetCoursesByIDs(ctx context.Context, req *pb.GetCoursesByIDsRequest) (*pb.GetCoursesByIDsResponse, error) {
	out, err := s.coursesService.GetCoursesByIDs(ctx, req.GetIds())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting courses by ids: %v", err)
	}
	return &pb.GetCoursesByIDsResponse{
		Courses: convertCoursesToPb(out),
	}, nil
}

func (s *server) GetAllCategories(ctx context.Context, req *pb.GetCategoriesRequest) (*pb.GetCategoriesResponse, error) {
	out, err := s.coursesService.GetAllCategories(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting all categories: %v", err)
	}
	return &pb.GetCategoriesResponse{
		Categories: convertCategoriesToPb(out),
	}, nil
}

func (s *server) GetContentByTopicID(ctx context.Context, req *pb.GetContentByTopicIDRequest) (*pb.GetContentByTopicIDResponse, error) {
	out, err := s.coursesService.GetContentByTopicID(ctx, req.GetCourseId(), req.GetTopicId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting content by topic id: %v", err)
	}
	return &pb.GetContentByTopicIDResponse{
		Content: convertContentToPb(out),
	}, nil
}

func (s *server) GetQuizzesByTopicID(ctx context.Context, req *pb.GetQuizzesByTopicIDRequest) (*pb.GetQuizzesByTopicIDResponse, error) {
	out, passed, err := s.coursesService.GetQuizzesByTopicID(
		ctx,
		req.UserId,
		req.CourseId,
		req.TopicId,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting quizzes by topic id: %v", err)
	}
	return &pb.GetQuizzesByTopicIDResponse{
		Quizzes: convertQuizzesToPb(out...),
		Passed:  passed,
	}, nil
}

func (s *server) GetTasksByTopicIDAndOrderNum(ctx context.Context, req *pb.GetTasksByTopicIDAndOrderNumRequest) (*pb.GetTasksByTopicIDAndOrderNumResponse, error) {
	out, _, err := s.coursesService.GetTasksByTopicIDAndOrderNum(
		ctx,
		req.CourseId,
		req.TopicId,
		int(req.Order),
		req.UserId,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting tasks by topic id and order num: %v", err)
	}
	return &pb.GetTasksByTopicIDAndOrderNumResponse{
		Task: convertPracticalTaskToPb(out),
	}, nil
}
