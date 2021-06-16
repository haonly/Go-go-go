package vin

type VinRepository interface {
	Save(vin string, vid string) (string, error)
	Get(vin string) (string, error)
	Delete(vin string) (string, error)
}
