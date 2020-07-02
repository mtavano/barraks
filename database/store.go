package database

import (
  "database/sql"
  "github.com/mtavano/barraks"
  _ "github.com/mattn/go-sqlite3"
)

const(
  insertInto = `INSERT INTO items (
    name,
    img_url,
    stock,
    unit,
    min_stock
  ) VALUES (?, ?, ?, ?, ?)`
  selectAllRecords = "SELECT * FROM items"
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

// Insert inserts an item into database and return an error if anything happens
func (st *Store) Insert(item barraks.Item) error {
  stmt, err := st.Prepare(insertInto)
  if err != nil {
    return err
  }
  defer stmt.Close()

  _, err = stmt.Exec(item.Name, item.ImgURL, item.Stock, item.Unit, item.MinStock)
  if err != nil {
    return err
  }

  return nil
}

// InsertScan inserts an item into database
func (st *Store) InsertScan(i *barraks.Item) error {
  err := st.Insert(*i)
  if err != nil {
    return err
  }

  return nil
}
