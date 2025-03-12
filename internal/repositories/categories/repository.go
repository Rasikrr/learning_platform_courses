package categories

import (
	"context"
	"github.com/Rasikrr/learning_platform_core/database"
	"github.com/Rasikrr/learning_platform_courses/internal/domain/entity"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/lib/pq"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*entity.Category, error)
	GetByIDs(ctx context.Context, ids []string) ([]*entity.Category, error)
	GetByID(ctx context.Context, id string) (*entity.Category, error)
}

type repository struct {
	db *database.Postgres
}

func NewRepository(db *database.Postgres) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetByID(ctx context.Context, id string) (*entity.Category, error) {
	var m model
	if err := pgxscan.Get(ctx, r.db, &m, getCategoriesByIDStmt, id); err != nil {
		return nil, err
	}
	return m.convert()
}

func (r *repository) GetByIDs(ctx context.Context, ids []string) ([]*entity.Category, error) {
	var mm models
	if err := pgxscan.Select(ctx, r.db, &mm, getCategoriesByIDsStmt, pq.Array(ids)); err != nil {
		return nil, err
	}
	return mm.convert()
}

func (r *repository) GetAll(ctx context.Context) ([]*entity.Category, error) {
	var mm models
	if err := pgxscan.Select(ctx, r.db, &mm, getAllCategoriesStmt); err != nil {
		return nil, err
	}
	return mm.convert()
}
