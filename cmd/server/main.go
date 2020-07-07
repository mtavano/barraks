package main

import (
  "net/http"
  "fmt"
  "os"
  "log"
	"github.com/gorilla/mux"
	"github.com/mtavano/barraks/cmd/server/handlers"
	"github.com/mtavano/barraks/database"
)

func main() {
  fmt.Println("[STARTING] Barraks")
	port := os.Getenv("PORT")
  dbFilePath := os.Getenv("DATABASE_FILE_PATH")
  log.Printf("database file paht: %s\n",dbFilePath)
  db, err := database.NewStore(dbFilePath)
  if err != nil {
    log.Fatalf(err.Error())
  }
  ctx := &handlers.Context{db}


	r := mux.NewRouter()

	addr := fmt.Sprintf(":%s", port)
  r.HandleFunc(handlers.HandleFuncWithCtx("/v1/items", ctx, handlers.CreateItem)).
    Methods(http.MethodPost, http.MethodOptions)
  r.HandleFunc(handlers.HandleFuncWithCtx("/v1/items", ctx, handlers.GetAll)).
    Methods(http.MethodGet, http.MethodOptions)
  r.HandleFunc(handlers.HandleFuncWithCtx("/v1/items/{id}", ctx, handlers.GetByID)).
    Methods(http.MethodGet, http.MethodOptions)
  r.HandleFunc(handlers.HandleFuncWithCtx("/v1/items/{id}", ctx, handlers.UpdateItem)).
    Methods(http.MethodPatch, http.MethodOptions)

	log.Printf("Server starting at %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
