package database

import (
  "fmt"
  "github.com/mtavano/barraks"
)


func (st *Store) SelectByID(id string) (*barraks.Item, error) {
  row := st.QueryRow(selectByID, id)
  i := new(barraks.Item)
  err := row.Scan(&i.ID, &i.Name, &i.ImgURL, &i.Stock, &i.Unit, &i.MinStock)
  if err != nil {
    fmt.Println(3, err)
    return nil, err
  }

  return i, nil

}

func (st *Store) SelectAll() ([]*barraks.Item, error) {
  items := make([]*barraks.Item, 0)
  rows, err := st.Query(selectAllRecords);
  if err != nil {
    return items, err
  }
  defer rows.Close()

  for rows.Next() {
    i := new(barraks.Item)
    if err := rows.Scan(&i.ID, &i.Name, &i.ImgURL, &i.Stock, &i.Unit, &i.MinStock); err != nil {
      return items, err
    }
    items = append(items, i)
  }

  return items, nil
}
