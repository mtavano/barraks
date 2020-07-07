package database

import (
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
)

type Store struct {
  *sql.DB
}

// NewStore ...
func NewStore(dbPath string) (*Store, error) {
  db, err := sql.Open("sqlite3", dbPath)
  if err != nil {
    return nil, err
  }

  return &Store{db}, nil
}

