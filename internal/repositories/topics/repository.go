package topics

import (
	"context"
	"github.com/Rasikrr/learning_platform_core/database"
	"github.com/Rasikrr/learning_platform_courses/internal/domain/entity"
	"github.com/georgysavva/scany/v2/pgxscan"
)

type Repository interface {
	GetByCourseID(ctx context.Context, id string) ([]*entity.Topic, error)
}

type repository struct {
	db *database.Postgres
}

func NewRepository(db *database.Postgres) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetByCourseID(ctx context.Context, ids string) ([]*entity.Topic, error) {
	var mm models
	if err := pgxscan.Select(ctx, r.db, &mm, getByCourseIDStmt, ids); err != nil {
		return nil, err
	}
	return mm.convert()
}
