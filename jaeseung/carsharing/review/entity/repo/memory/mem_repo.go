package memory

import (
	"context"
	entity2 "github.com/haonly/Go-go-go/jaeseung/carsharing/review/entity"
)

type MemRepository struct {
	repo map[string]entity2.Review
}

func New() *MemRepository {
	m := make(map[string]entity2.Review)
	return &MemRepository{repo: m}
}

func (m *MemRepository) Save(ctx context.Context, r *entity2.Review) error {
	m.repo[r.Id] = *r
	return nil
}
