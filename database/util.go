package database

import (
	"os"
	"path/filepath"
)

func setupPath(path string) error {
	// Check if the path exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Create the path if it does not exist
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// Check if the path contains the database folder
	dbFolderPath := filepath.Join(path, "database")
	if _, err := os.Stat(dbFolderPath); os.IsNotExist(err) {
		// Create the database folder if it does not exist
		err := os.Mkdir(dbFolderPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// Check if the path contains the manifest folder
	manifestFolderPath := filepath.Join(path, "manifest")
	if _, err := os.Stat(manifestFolderPath); os.IsNotExist(err) {
		// Create the manifest folder if it does not exist
		err := os.Mkdir(manifestFolderPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

func ReadManifest(path string) (*Manifest, error) {
	// Read the manifest file
	serializedManifest, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var manifest Manifest
	err = manifest.Deserialize(serializedManifest)
	if err != nil {
		return nil, err
	}

	return &manifest, nil
}
