package database

import (
	"os"
	"path/filepath"
)

func New(path string) (*DB, error) {
	err := setupPath(path)
	if err != nil {
		return nil, err
	}

	return &DB{
		base_path: path,
	}, nil
}

type DB struct {
	base_path string
}

func (db *DB) CreateDB(manifest *Manifest) error {
	// Check if the manifest is exist in the database
	manifest, err := NewManifest(manifest.Name)
	if err != nil {
		return err
	}

	// Serialize the manifest
	serializedManifest, err := manifest.Serialize()
	if err != nil {
		return err
	}

	// Write the serialized manifest to a file
	manifestPath := filepath.Join(db.base_path, "manifest", manifest.ID.String()+".json")
	err = os.WriteFile(manifestPath, serializedManifest, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) RemoveDB(manifest *Manifest) error {
    // Remove the manifest file
    manifestPath := filepath.Join(db.base_path, "manifest", manifest.ID.String()+".json")
    err := os.Remove(manifestPath)
    if err != nil {
        return err
    }

    return nil
}

func (db *DB) ListDBs() ([]Manifest, error) {
	// List all the databases in the database folder
	manifestFolderPath := filepath.Join(db.base_path, "manifest")
	files, err := os.ReadDir(manifestFolderPath)
	if err != nil {
		return nil, err
	}

	var databases []string
	for _, file := range files {
		databases = append(databases, file.Name())
	}

	// Read the manifest files
	var manifests []Manifest
	for _, database := range databases {
		manifest, err := ReadManifest(filepath.Join(manifestFolderPath, database))
		if err != nil {
			return nil, err
		}

		manifests = append(manifests, *manifest)
	}

	return manifests, nil
}
