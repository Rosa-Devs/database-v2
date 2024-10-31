package database

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/google/uuid"
)

// GenerateManifest generates a new manifest with AES-256-GCM encryption and checksum
func NewManifest(name string) (*Manifest, error) {
	// Generate a new UUID for the manifest
	manifestID, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate UUID: %v", err)
	}

	// Generate a random AES-256-GCM secret key
	secret, err := generateSecretKey(32) // 32 bytes for AES-256
	if err != nil {
		return nil, fmt.Errorf("failed to generate secret: %v", err)
	}

    randevuz, err := generateSecretKey(32)
    if err != nil {
        return nil, fmt.Errorf("failed to generate secret: %v", err)
    }

	// Create the manifest
	manifest := &Manifest{
		ID:                  manifestID,
		Name:                name,
		EncryptionAlgorithm: "AES-256-GCM",
		Secret:              hex.EncodeToString(secret), // Convert secret to hexadecimal string
        Randevuz: hex.EncodeToString(randevuz),
		CreatedAt:           time.Now().UTC(),
	}

	return manifest, nil
}

// generateSecretKey generates a random secret key of the specified size (in bytes)
func generateSecretKey(size int) ([]byte, error) {
	key := make([]byte, size)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return nil, fmt.Errorf("failed to generate random key: %v", err)
	}
	return key, nil
}

type Manifest struct {
	ID                  uuid.UUID `json:"id"`
	Name                string    `json:"name"`
	EncryptionAlgorithm string    `json:"encryption_algorithm"`
	Secret              string    `json:"secret"`
    Randevuz            string    `json:"randevuz"`
	CreatedAt           time.Time `json:"created_at"`
}

func (m *Manifest) Serialize() ([]byte, error) {
	// JSON serialization
	bytes, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
func (m *Manifest) Deserialize(data []byte) error {
	// JSON deserialization
	err := json.Unmarshal(data, m)
	if err != nil {
		return err
	}
	return nil
}

type Basemodel struct {
	ID       uuid.UUID `json:"id"`
	PoolName string    `json:"pool_name"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy string     `json:"updated_by"`
	DeleteAt  *time.Time `json:"delete_at"`
}
