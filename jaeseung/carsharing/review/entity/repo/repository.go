package repo

import (
	"context"
	entity2 "github.com/haonly/Go-go-go/jaeseung/carsharing/review/entity"
)

type ReviewRepository interface {
	Save(ctx context.Context, r *entity2.Review) error
}
