package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Rosa-Devs/database-v2/database"
)

func main() {
	// Get the working directory
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	db, err := database.New(workingDir + "/temp/db1")
	if err != nil {
		panic(err)
	}

	manifest, err := database.NewManifest("my-manifest")
	if err != nil {
		panic(err)
	}

	err = db.CreateDB(manifest)
	if err != nil {
		panic(err)
	}

	manifests, err := db.ListDBs()
	if err != nil {
		panic(err)
	}

	json_data, err := json.Marshal(manifests)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(json_data))

    // Remove the manifest
    err = db.RemoveDB(manifest)
    if err != nil {
        panic(err)
    }
}
