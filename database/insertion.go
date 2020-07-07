package database

import (
  "fmt"
  "strings"
  "github.com/mtavano/barraks"
  "github.com/google/uuid"
)
// Insert inserts an item into database and return an error if anything happens
func (st *Store) Insert(i barraks.Item) (*string, error) {
  stmt, err := st.Prepare(insertInto)
  defer stmt.Close()
  if err != nil {
    fmt.Println(1, err)
    return nil, err
  }

  sid := uuid.New().String()
  name := strings.ToLower(i.Name)
  unit := strings.ToLower(i.Unit)
  _, err = stmt.Exec(sid, name, i.ImgURL, i.Stock, unit, i.MinStock)
  if err != nil {
    fmt.Println(2, err)
    return nil, err
  }

  return &sid, nil
}

// InsertScan inserts an item into database
func (st *Store) InsertAndReturn(i barraks.Item) (*barraks.Item, error) {
  id, err := st.Insert(i)
  if err != nil {
    return nil, err
  }

  item, err := st.SelectByID(*id)
  if err != nil {
    return nil, err
  }

  return item, nil
}

func (st *Store) UpdateScan(i *barraks.Item) error {
  tx, err := st.Begin()
  if err != nil {
    return err
  }

  _, err = tx.Exec(updateRecord, i.Name, i.ImgURL, i.Stock, i.Unit, i.MinStock, i.ID)
  if err != nil {
    tx.Rollback()
    return err
  }

  row := tx.QueryRow(selectByID, i.ID)
  err = row.Scan(&i.ID, &i.Name, &i.ImgURL, &i.Stock, &i.Unit, &i.MinStock)
  if err != nil {
    tx.Rollback()
    fmt.Println(4, err)
    return err
  }

  err = tx.Commit()
  if err != nil {
    return err
  }

  return nil
}

