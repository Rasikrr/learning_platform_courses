package grpc

import (
	"github.com/Rasikrr/learning_platform_core/grpc/converters"
	"github.com/Rasikrr/learning_platform_courses/internal/domain/entity"
	pb "github.com/Rasikrr/learning_platform_courses/pkg/api/proto/courses"
	"github.com/samber/lo"
)

func convertCreateCourseReq(req *pb.CreateCourseRequest) *entity.CreateCourseParams {
	return &entity.CreateCourseParams{
		Title:       req.Title,
		ImageURL:    req.ImageUrl,
		CategoryID:  req.CategoryId,
		Description: req.Description,
		CreatedBy:   req.CreatedBy,
	}
}

func convertGetCoursesByParamsReq(req *pb.GetCoursesByParamsRequest) *entity.GetCoursesParams {
	return &entity.GetCoursesParams{
		Limit:         int(req.GetLimit()),
		Offset:        int(req.GetOffset()),
		CategoriesIDs: req.GetCategoryIds(),
	}
}

func convertCoursesToPb(courses []*entity.Course) []*pb.Course {
	if len(courses) == 0 {
		return nil
	}
	return lo.Map(courses, func(course *entity.Course, _ int) *pb.Course {
		return convertCourseToPb(course)
	})
}

func convertCourseToPb(course *entity.Course) *pb.Course {
	if course == nil {
		return nil
	}
	return &pb.Course{
		Id:          course.ID,
		Title:       course.Title,
		ImageUrl:    course.ImageURL,
		Category:    convertCategoryToPb(&course.Category),
		Description: course.Description,
		Topics:      convertTopicsToPb(course.Topics...),
		CreatedAt:   converters.ConvertToTimestampPb(&course.CreatedAt),
		UpdatedAt:   converters.ConvertToTimestampPb(&course.UpdatedAt),
	}
}

func convertCategoriesToPb(categories []*entity.Category) []*pb.Category {
	return lo.Map(categories, func(item *entity.Category, _ int) *pb.Category {
		return convertCategoryToPb(item)
	})
}

func convertCategoryToPb(category *entity.Category) *pb.Category {
	if category == nil {
		return nil
	}
	return &pb.Category{
		Id:        category.ID,
		Name:      category.Name,
		CreatedBy: category.CreatedBy,
		CreatedAt: converters.ConvertToTimestampPb(&category.CreatedAt),
		UpdatedAt: converters.ConvertToTimestampPb(&category.UpdatedAt),
	}
}

func convertTopicsToPb(topics ...*entity.Topic) []*pb.Topic {
	if len(topics) == 0 {
		return nil
	}
	return lo.Map(topics, func(topic *entity.Topic, _ int) *pb.Topic {
		return &pb.Topic{
			Id:             topic.ID,
			CourseId:       topic.CourseID,
			Title:          topic.Title,
			Description:    topic.Description,
			Content:        convertContentToPb(topic.Content),
			Quizzes:        convertQuizzesToPb(topic.Quizzes...),
			PracticalTasks: convertPracticalTasksToPb(topic.PracticalTasks...),
			OrderNumber:    int32(topic.OrderNumber),
			CreatedAt:      converters.ConvertToTimestampPb(&topic.CreatedAt),
			UpdatedAt:      converters.ConvertToTimestampPb(&topic.UpdatedAt),
		}
	})
}

func convertContentToPb(content *entity.TopicContent) *pb.TopicContent {
	if content == nil {
		return nil
	}
	return &pb.TopicContent{
		Id:                  content.ID,
		TopicId:             content.TopicID,
		Content:             content.Content,
		AdditionalResources: content.AdditionalResources,
		VideoUrls:           content.VideoURLs,
		ImageUrls:           content.ImageURLs,
		CreatedAt:           converters.ConvertToTimestampPb(&content.CreatedAt),
		UpdatedAt:           converters.ConvertToTimestampPb(&content.UpdatedAt),
	}
}

func convertQuizzesToPb(quizzes ...*entity.Quiz) []*pb.Quiz {
	if len(quizzes) == 0 {
		return nil
	}
	return lo.Map(quizzes, func(quiz *entity.Quiz, _ int) *pb.Quiz {
		return &pb.Quiz{
			Id:             quiz.ID,
			TopicId:        quiz.TopicID,
			Question:       quiz.Question,
			Options:        quiz.Options,
			CorrectAnswers: quiz.CorrectAnswers,
			MultipleChoice: quiz.MultipleChoice,
			CreatedAt:      converters.ConvertToTimestampPb(&quiz.CreatedAt),
			UpdatedAt:      converters.ConvertToTimestampPb(&quiz.UpdatedAt),
		}
	})
}

func convertPracticalTasksToPb(tasks ...*entity.PracticalTask) []*pb.PracticalTask {
	if len(tasks) == 0 {
		return nil
	}
	return lo.Map(tasks, func(task *entity.PracticalTask, _ int) *pb.PracticalTask {
		return convertPracticalTaskToPb(task)
	})
}

func convertPracticalTaskToPb(task *entity.PracticalTask) *pb.PracticalTask {
	if task == nil {
		return nil
	}
	return &pb.PracticalTask{
		Id:              task.ID,
		TopicId:         task.TopicID,
		Description:     task.Description,
		DifficultyLevel: task.DifficultyLevel.String(),
		StarterCode:     task.StarterCode,
		ExpectedOutput:  task.ExpectedOutput,
		OrderNumber:     int32(task.OrderNumber),
		CreatedAt:       converters.ConvertToTimestampPb(&task.CreatedAt),
		UpdatedAt:       converters.ConvertToTimestampPb(&task.UpdatedAt),
		TestCases:       task.TestCases,
		Language:        task.Language.String(),
	}
}
