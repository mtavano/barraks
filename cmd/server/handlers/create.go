package handlers

import (
  "net/http"
  "github.com/mtavano/barraks"
)

// CreateItem ....
func CreateItem(ctx *Context, r *http.Request) (*Response, int) {
    payload := &struct{
      Item barraks.Item `json:"item"`
    }{}

    if err := unmarshalJSON(r.Body, payload); err != nil {
      return &Response{ Error: err }, http.StatusInternalServerError
    }

    item, err := ctx.InsertAndReturn(payload.Item)
    // assuming that the insertion went wrong because of any
    // trouble with missing or bad data 
    if err != nil {
      return &Response{ Error: err }, http.StatusBadRequest
    }

    return &Response{ Data: item }, http.StatusCreated
}

