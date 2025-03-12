package categories

import (
	"github.com/Rasikrr/learning_platform_courses/internal/domain/entity"
	"time"
)

type model struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	CreatedBy string    `db:"created_by"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type models []model

// nolint: unparam
func (m model) convert() (*entity.Category, error) {
	return &entity.Category{
		ID:        m.ID,
		Name:      m.Name,
		CreatedBy: m.CreatedBy,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}, nil
}

func (mm models) convert() ([]*entity.Category, error) {
	out := make([]*entity.Category, len(mm))
	for i, m := range mm {
		res, err := m.convert()
		if err != nil {
			return nil, err
		}
		out[i] = res
	}
	return out, nil
}
