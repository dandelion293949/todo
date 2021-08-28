package main

import (
  "net/http"
  "encoding/json"
  "time"
  "fmt"
  "log"
  "io"
  "bytes"

  "github.com/dandelion293949/todo/domains"
)

var (
  repo = domains.New()
)

func main() {

  http.HandleFunc("/todos/", handler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
  if r.Method == http.MethodGet {
    id := r.URL.Path[len("/todos/"):]
    if id == "" {
      todos, err := repo.GetAll()
      if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "invalid request. err = %v", err)
        return
      }
      b, err := json.Marshal(todos)
      if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "invalid request. err = %v", err)
        return
      }
      fmt.Fprintf(w, "%s", b)
    } else {
      todo, err := repo.Get(domains.TodoID(id))
      if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "invalid request. err = %v", err)
        return
      }
      b, err := json.Marshal(*todo)
      if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "invalid request. err = %v", err)
        return
      }
      fmt.Fprintf(w, "%s", b)
    }
  }

  body := r.Body
  defer body.Close()

  buf := new(bytes.Buffer)
  io.Copy(buf, body)

  var todo domains.Todo
  json.Unmarshal(buf.Bytes(), &todo)

  if r.Method == http.MethodPost {
    t, err := repo.Create(&todo, time.Now())
    if err != nil {
      w.WriteHeader(http.StatusBadRequest)
      fmt.Fprintf(w, "invalid request. err = %v", err)
      return
    }

    b, err := json.Marshal(*t)
    if err != nil {
      w.WriteHeader(http.StatusBadRequest)
      fmt.Fprintf(w, "invalid request. err = %v", err)
      return
    }
     w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "%s", b)
  }

  if r.Method == http.MethodPut {
    t, err := repo.Update(&todo, time.Now())
    if err != nil {
      w.WriteHeader(http.StatusBadRequest)
      fmt.Fprintf(w, "invalid request. err = %v", err)
      return
    }

    b, err := json.Marshal(*t)
    if err != nil {
      w.WriteHeader(http.StatusBadRequest)
      fmt.Fprintf(w, "invalid request. err = %v", err)
      return
    }
     w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "%s", b)
  }

  if r.Method == http.MethodDelete {
    t, err := repo.Delete(&todo)
    if err != nil {
      w.WriteHeader(http.StatusBadRequest)
      fmt.Fprintf(w, "invalid request. err = %v", err)
      return
    }

    b, err := json.Marshal(*t)
    if err != nil {
      w.WriteHeader(http.StatusBadRequest)
      fmt.Fprintf(w, "invalid request. err = %v", err)
      return
    }
     w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "%s", b)
  }
}

