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
  url = "/todos/"
)

func main() {

  http.HandleFunc(url, handler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func get(path string) []byte {
  id := path[len(url):]
  if id == "" {
    todos, err := repo.GetAll()
    if err != nil {
      return []byte(`{ "error": "could not get todo"}`)
    }
    b, err := json.Marshal(todos)
    if err != nil {
      return []byte(`{ "error": "could not converted json"}`)
    }
    return b
  } else {
    todo, err := repo.Get(domains.TodoID(id))
    if err != nil {
      return []byte(`{ "error": "could not get todos"}`)
    }
    b, err := json.Marshal(*todo)
    if err != nil {
      return []byte(`{ "error": "could not converted json"}`)
    }
    return b
  }
}

func create(todo domains.Todo) []byte {
  t, err := repo.Create(&todo, time.Now())
  if err != nil {
    return []byte(`{ "error": "invalid request" }`)
  }

  b, err := json.Marshal(*t)
  if err != nil {
    return []byte(`{ "error": "could not converted json"}`)
  }

  return b
}

func update(todo domains.Todo) []byte {
  t, err := repo.Update(&todo, time.Now())
  if err != nil {
    return []byte(`{ "error": "invalid request" }`)
  }

  b, err := json.Marshal(*t)
  if err != nil {
    return []byte(`{ "error": "could not converted json"}`)
  }
  return b
}

func delete(todo domains.Todo) []byte {
  t, err := repo.Delete(&todo)
  if err != nil {
    return []byte(`{ "error": "invalid request" }`)
  }

  b, err := json.Marshal(*t)
  if err != nil {
    return []byte(`{ "error": "could not converted json"}`)
  }
  return b
}

func handler(w http.ResponseWriter, r *http.Request) {
  if r.Method == http.MethodGet {
    fmt.Fprintf(w, "%s", get(r.URL.Path))
  }

  body := r.Body
  defer body.Close()

  buf := new(bytes.Buffer)
  io.Copy(buf, body)

  var todo domains.Todo
  json.Unmarshal(buf.Bytes(), &todo)

  if r.Method == http.MethodPost {
    fmt.Fprintf(w, "%s", create(todo))
  }

  if r.Method == http.MethodPut {
    fmt.Fprintf(w, "%s", update(todo))
  }

  if r.Method == http.MethodDelete {
    fmt.Fprintf(w, "%s", delete(todo))
  }
}

