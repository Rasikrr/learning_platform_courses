package app

import (
	"context"
	"github.com/Rasikrr/learning_platform_core/application"
	coursesC "github.com/Rasikrr/learning_platform_courses/internal/cache/courses"
	"github.com/Rasikrr/learning_platform_courses/internal/ports/grpc"
	categoriesR "github.com/Rasikrr/learning_platform_courses/internal/repositories/categories"
	contentR "github.com/Rasikrr/learning_platform_courses/internal/repositories/content"
	coursesR "github.com/Rasikrr/learning_platform_courses/internal/repositories/courses"
	quizzesR "github.com/Rasikrr/learning_platform_courses/internal/repositories/quizzes"
	quizzesSubmissionsR "github.com/Rasikrr/learning_platform_courses/internal/repositories/quizzes_submissions"
	tasksR "github.com/Rasikrr/learning_platform_courses/internal/repositories/tasks"
	tasksSubmissionsR "github.com/Rasikrr/learning_platform_courses/internal/repositories/tasks_submissions"
	topicsR "github.com/Rasikrr/learning_platform_courses/internal/repositories/topics"
	coursesS "github.com/Rasikrr/learning_platform_courses/internal/services/courses"
)

type App struct {
	*application.App
	quizzesRepository           quizzesR.Repository
	topicsRepository            topicsR.Repository
	coursesRepository           coursesR.Repository
	categoriesRepository        categoriesR.Repository
	tasksRepository             tasksR.Repository
	contentRepository           contentR.Repository
	quizzesSubmissionRepository quizzesSubmissionsR.Repository
	taskSubmissionRepository    tasksSubmissionsR.Repository
	coursesCache                coursesC.Cache
	coursesService              coursesS.Service
}

func NewApp(ctx context.Context, name string) (*App, error) {
	app := &App{
		App: application.NewApp(ctx, name),
	}
	if err := app.Init(ctx); err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Init(ctx context.Context) error {
	for _, init := range []func(context.Context) error{
		a.initRepositories,
		a.initCaches,
		a.initClients,
		a.initServices,
		a.initGRPCServer,
	} {
		if err := init(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initRepositories(_ context.Context) error {
	a.quizzesRepository = quizzesR.NewRepository(a.Postgres())
	a.topicsRepository = topicsR.NewRepository(a.Postgres())
	a.categoriesRepository = categoriesR.NewRepository(a.Postgres())
	a.coursesRepository = coursesR.NewRepository(a.Postgres())
	a.tasksRepository = tasksR.NewRepository(a.Postgres())
	a.contentRepository = contentR.NewRepository(a.Postgres())
	a.quizzesSubmissionRepository = quizzesSubmissionsR.NewRepository(a.Postgres())
	a.taskSubmissionRepository = tasksSubmissionsR.NewRepository(a.Postgres())
	return nil
}

func (a *App) initCaches(ctx context.Context) error {
	a.coursesCache = coursesC.NewCache(a.Redis())
	return nil
}

func (a *App) initClients(ctx context.Context) error {
	return nil
}

func (a *App) initServices(_ context.Context) error {
	a.coursesService = coursesS.NewService(
		a.coursesRepository,
		a.categoriesRepository,
		a.topicsRepository,
		a.quizzesRepository,
		a.tasksRepository,
		a.contentRepository,
		a.quizzesSubmissionRepository,
		a.taskSubmissionRepository,
		a.coursesCache,
	)
	return nil
}

func (a *App) initGRPCServer(_ context.Context) error {
	grpc.NewServer(
		a.GrpcServer().Srv(),
		a.coursesService,
	)
	return nil
}
