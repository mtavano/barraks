package main

import (
  "io"
  "io/ioutil"
  "os"
  "log"
  "fmt"
  "encoding/json"
  "github.com/mtavano/barraks"
  "github.com/mtavano/barraks/database"

  _ "github.com/mattn/go-sqlite3"
)

func main() {
  items := getSeedData("./data.json")
  store, err := database.NewStore("../server/barraks.db")
  if err != nil {
    log.Fatal(err)
  }

  seedDatabase(store, items)
}

func seedDatabase(db *database.Store, items []barraks.Item) {
  for _, item := range items {
    err := db.Insert(item)
    if err != nil {
      log.Fatalf("[FAIL] INSERT\n %s", err.Error())
    }
    fmt.Printf("[OK] INSERT")
    printStruct(item)
  }
}

func getSeedData(fileLocation string) []barraks.Item {
  fileReader, err := os.Open(fileLocation)
  if err != nil {
    log.Fatal(err)
  }
  defer fileReader.Close()

  items := make([]barraks.Item, 0)
  err = unmarshalJSON(fileReader, &items)
  if err != nil {
    log.Fatal(err)
  }

  return items
}

func printStruct(it barraks.Item) {
  fmt.Printf("%+v\n", it)
}


func unmarshalJSON(r io.Reader, v interface{}) error {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}
