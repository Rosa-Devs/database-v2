package worker

import "github.com/Rosa-Devs/database-v2/database"


type Worker struct {
    // Database
    DB *database.DB
}

func NewWorker(db *database.DB) (*Worker, error) {
    return &Worker{
        DB: db,
    }, nil
}

func (w *Worker) Start() error {
    // Start the worker
    return nil
}
