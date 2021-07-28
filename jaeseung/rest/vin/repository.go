package vin

import "context"

type VinRepository interface {
	Save(ctx context.Context, vin string, vid string) (string, error)
	Get(ctx context.Context, vin string) (string, error)
	Update(ctx context.Context, vin string, vid string) error
	Delete(ctx context.Context, vin string) (string, error)
}
