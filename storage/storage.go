package storage

import "github.com/Rosa-Devs/database-v2/database"

type Storage struct {
	DB       *database.DB
	Manifest *database.Manifest

	// Add more fields here
}

func NewStorage(db *database.DB, manifest *database.Manifest) (*Storage, error) {
	return &Storage{
		DB:       db,
		Manifest: manifest,
	}, nil
}

func (s *Storage) CreatePool(poolName string) error {
	// Create a new pool in the database
	return nil
}
