package usecase

import (
	"context"
	"github.com/haonly/Go-go-go/jaeseung/carsharing/review/entity"
	"github.com/haonly/Go-go-go/jaeseung/carsharing/review/entity/repo"
	"github.com/sirupsen/logrus"
)

type ReviewService struct {
	log  *logrus.Entry
	repo repo.ReviewRepository
}

func NewReview(repo repo.ReviewRepository) *ReviewService {
	return &ReviewService{
		log:  logrus.WithFields(logrus.Fields{"tag": "review"}),
		repo: repo,
	}
}

type CreateReviewResDTO struct {
}

type CreateReviewReqDTO struct {
	Ctx            context.Context
	CreatorId      string
	ReviewTargetId string
	Text           string
	Rating         float32
}

func (c *CreateReviewReqDTO) Context() context.Context {
	return c.Ctx
}

func (r *ReviewService) CreateReview(req CreateReviewReqDTO) (*CreateReviewResDTO, error) {
	newReview := entity.NewReview(req.CreatorId, req.ReviewTargetId)
	if err := newReview.Validate(); err != nil {
		return nil, err
	}
	if err := r.repo.Save(req.Context(), newReview); err != nil {
		return nil, err
	}
	return &CreateReviewResDTO{}, nil
}
