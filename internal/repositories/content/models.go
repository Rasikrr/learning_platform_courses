package content

import (
	"github.com/Rasikrr/learning_platform_courses/internal/domain/entity"
	"time"
)

type model struct {
	ID                  string    `db:"id"`
	TopicID             string    `db:"topic_id"`
	Content             string    `db:"content"`
	AdditionalResources []string  `db:"additional_resources"`
	VideoURLs           []string  `db:"video_urls"`
	ImageURLs           []string  `db:"image_urls"`
	CreatedAt           time.Time `db:"created_at"`
	UpdatedAt           time.Time `db:"updated_at"`
}

type models []model

// nolint
func (m model) convert() (*entity.TopicContent, error) {
	return &entity.TopicContent{
		ID:                  m.ID,
		TopicID:             m.TopicID,
		Content:             m.Content,
		AdditionalResources: m.AdditionalResources,
		VideoURLs:           m.VideoURLs,
		ImageURLs:           m.ImageURLs,
		CreatedAt:           m.CreatedAt,
		UpdatedAt:           m.UpdatedAt,
	}, nil
}

func (mm models) convert() ([]*entity.TopicContent, error) {
	out := make([]*entity.TopicContent, len(mm))
	for i, m := range mm {
		res, err := m.convert()
		if err != nil {
			return nil, err
		}
		out[i] = res
	}
	return out, nil
}
