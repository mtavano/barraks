package handlers

import (
  "net/http"
  "github.com/gorilla/mux"
)

// GetByName ...
func GetByID(ctx *Context, r *http.Request) (*Response, int) {
  vars := mux.Vars(r)
  id := vars["id"]
  item, err := ctx.SelectByID(id)
  if err != nil {
    return &Response{Error: err}, http.StatusBadRequest
  }

  return &Response{ Data: item }, http.StatusOK
}

func GetAll(ctx *Context, r *http.Request) (*Response, int) {
  items, err := ctx.SelectAll()
  if err != nil {
    return &Response{Error: err}, http.StatusInternalServerError
  }

  return &Response{ Data: items }, http.StatusOK
}
