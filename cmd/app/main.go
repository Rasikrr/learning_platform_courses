package main

import (
	"context"
	"github.com/Rasikrr/learning_platform_courses/internal/app"
	"log"
)

const (
	appName = "courses"
)

func main() {
	ctx := context.Background()
	app, err := app.NewApp(ctx, appName)
	if err != nil {
		log.Fatalf("error in init app: %v", err)
	}
	if err := app.Start(ctx); err != nil {
		log.Fatalf("error in start app: %v", err)
	}
}
