package memory

import (
	"context"
	"errors"
)

type MemRepository struct {
	repo map[string]string
}

func New() *MemRepository {
	m := make(map[string]string)
	return &MemRepository{repo: m}
}

func (m *MemRepository) Save(ctx context.Context, vin string, vid string) (string, error) {
	m.repo[vin] = vid
	return vid, nil
}

func (m *MemRepository) Get(ctx context.Context, vin string) (string, error) {
	vid, ok := m.repo[vin]
	if !ok {
		return "", errors.New("Not Found")
	}
	return vid, nil
}

func (m *MemRepository) Update(ctx context.Context, vin string, vid string) error {
	_, err := m.Get(ctx, vin)
	if err != nil {
		return err
	}

	_, err = m.Save(ctx, vin, vid)
	if err != nil {
		return err
	}
	return nil
}

func (m *MemRepository) Delete(ctx context.Context, vin string) (string, error) {
	vid, ok := m.repo[vin]
	if !ok {
		return "", errors.New("Not Found")
	}
	delete(m.repo, vin)
	return vid, nil
}
