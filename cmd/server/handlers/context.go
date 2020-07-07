package handlers

import (
  "io"
  "io/ioutil"
  "net/http"
  "log"
  "encoding/json"
  "github.com/mtavano/barraks/database"
)

type FuncHandler func(*Context, *http.Request)(*Response, int)
type HttpHandler func(http.ResponseWriter, *http.Request)

type Context struct {
  *database.Store
}

type Response struct {
  Data interface{}
  Meta interface{}
  Error error
}


func (r *Response) Build() map[string]interface{} {
  m := make(map[string]interface{})

  if r.Data != nil {
    m["data"] = r.Data
  }

  if r.Meta != nil {
    m["meta"] = r.Meta
  }

  if r.Error != nil {
    m["meta"] = r.Error.Error()
  }

  return m
}

func HandleFuncWithCtx(path string, ctx *Context, f FuncHandler) (string, HttpHandler) {
  return path, func(w http.ResponseWriter, r *http.Request) {
    log.Printf("START [%s] %s", r.Method, path)
    res, status := f(ctx, r)
    if res.Error != nil {
      log.Printf("[ERROR] %+v\n", res.Error)
    }
    log.Printf("END [%s] %s STATUS: %d\n", r.Method, path, status)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)

    if err := json.NewEncoder(w).Encode(res.Build()); err != nil {
      log.Printf("%+v\n", err)
      http.Error(w, err.Error(), status)
    }
  }
}
func unmarshalJSON(r io.Reader, v interface{}) error {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}
