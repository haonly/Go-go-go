package memory

import "errors"

type MemRepository struct {
	repo map[string]string
}

func New() *MemRepository {
	m := make(map[string]string)
	return &MemRepository{repo: m}
}

func (m *MemRepository) Save(vin string, vid string) (string, error) {
	m.repo[vin] = vid
	return vid, nil
}

func (m *MemRepository) Get(vin string) (string, error) {
	vid, ok := m.repo[vin]
	if !ok {
		return "", errors.New("Not Found")
	}
	return vid, nil
}

func (m *MemRepository) Update(vin string, vid string) error {
	_, err := m.Get(vin)
	if err != nil {
		return err
	}

	_, err = m.Save(vin, vid)
	if err != nil {
		return err
	}
	return nil
}

func (m *MemRepository) Delete(vin string) (string, error) {
	vid, ok := m.repo[vin]
	if !ok {
		return "", errors.New("Not Found")
	}
	delete(m.repo, vin)
	return vid, nil
}
