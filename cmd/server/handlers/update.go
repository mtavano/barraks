package handlers

import (
  "net/http"
  "github.com/gorilla/mux"
  "github.com/mtavano/barraks"
)

func UpdateItem(ctx *Context, r *http.Request) (*Response, int) {
  payload := &struct{
    Item barraks.Item `json:"item"`
  }{}
  vars := mux.Vars(r)

   if err := unmarshalJSON(r.Body, payload); err != nil {
    return &Response{ Error: err }, http.StatusInternalServerError
  }

  payload.Item.ID = vars["id"]
  if err := ctx.UpdateScan(&payload.Item); err != nil {
    return &Response{ Error: err }, http.StatusBadRequest
  }

  return &Response{ Data: payload.Item }, http.StatusOK
}
