package entity

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Review struct {
	Id             string
	CreatorId      string
	ReviewTargetId string
	CreateTime     int64
	Text           string
	Rating         float32
}

func NewReview(creatorId, reviewTargetId string) *Review {
	return &Review{
		Id:             uuid.New().String(),
		CreatorId:      creatorId,
		ReviewTargetId: reviewTargetId,
		CreateTime:     time.Now().UTC().Unix(),
	}
}

func (r *Review) WithText(text string) *Review {
	r.Text = text
	return r
}

func (r *Review) WithRating(rating float32) *Review {
	r.Rating = rating
	return r
}

func (r *Review) Validate() error {
	if r.Rating <= 0 {
		return errors.New("Review Entity validation error")
	}
	return nil
}
