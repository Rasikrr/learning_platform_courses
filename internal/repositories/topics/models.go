package topics

import (
	"github.com/Rasikrr/learning_platform_courses/internal/domain/entity"
	"time"
)

type model struct {
	ID          string
	CourseID    string
	Title       string
	Description string
	OrderNumber int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type models []model

// nolint
func (m model) convert() (*entity.Topic, error) {
	return &entity.Topic{
		ID:          m.ID,
		CourseID:    m.CourseID,
		Title:       m.Title,
		Description: m.Description,
		OrderNumber: m.OrderNumber,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}, nil
}

// nolint
func (mm models) convert() ([]*entity.Topic, error) {
	out := make([]*entity.Topic, len(mm))
	for i, m := range mm {
		res, err := m.convert()
		if err != nil {
			return nil, err
		}
		out[i] = res
	}
	return out, nil
}
