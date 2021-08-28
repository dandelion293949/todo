package domains_test

import (
  "time"
  "testing"

  "github.com/dandelion293949/todo/domains"
)

func TestTodoCreate (t *testing.T) {
  type args struct {
    text string
    now time.Time
  }
  type want struct {
    text string
    createdAt time.Time
    updatedAt time.Time
  }
  tests := []struct {
    name string
    args args
    want want
  }{
    { name: "text_jp", args: args{ text: "明日から本気出す", now: time.Date(2021, 4, 1, 10, 0, 0, 0, time.UTC) }, want: want{ text: "明日から本気出す", createdAt: time.Date(2021, 4, 1, 10, 0, 0, 0, time.UTC), updatedAt: time.Time{} } },
    { name: "text_en", args: args{ text: "create todo", now: time.Date(2021, 4, 1, 10, 0, 0, 0, time.UTC) }, want: want{ text: "create todo", createdAt: time.Date(2021, 4, 1, 10, 0, 0, 0, time.UTC), updatedAt: time.Time{} } },
    { name: "symbol", args: args{ text: "`-=[]\\;',./~!@#$%^&*()_+{}|:\"<>?", now: time.Date(2021, 4, 1, 10, 0, 0, 0, time.UTC) }, want: want{ text: "`-=[]\\;',./~!@#$%^&*()_+{}|:\"<>?", createdAt: time.Date(2021, 4, 1, 10, 0, 0, 0, time.UTC), updatedAt: time.Time{} } },
  }

  repo := domains.New()
  todos, err := repo.GetAll()
  if err != nil {
    t.Fatalf("GetAll failed")
  }
  if len(todos) != 0 {
    t.Fatalf("repo = %v, want = %v", todos, []domains.Todo{})
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      todo, err := repo.Create(&domains.Todo{Text: tt.args.text}, tt.args.now)
      if err != nil {
        t.Fatalf("err = %v", err)
      }
      if todo.Id == "" {
        t.Fatalf("todo.Id = %v", todo.Id)
      }
      if todo.Text != tt.want.text {
        t.Fatalf("todo.Text = %v, want = %v", todo.Text, tt.want.text)
      }
      if todo.CreatedAt != tt.want.createdAt {
        t.Fatalf("todo.CreatedAt = %v, want = %v", todo.Text, tt.want.text)
      }
      if todo.UpdatedAt != tt.want.updatedAt {
        t.Fatalf("todo.UpdatedAt = %v, want = %v", todo.Text, tt.want.text)
      }
    })
  }
}

func TestTodoUpdate (t *testing.T) {
  repo := domains.New()
  todos, err := repo.GetAll()
  if err != nil {
    t.Fatalf("GetAll failed")
  }
  if len(todos) != 0 {
    t.Fatalf("repo = %v, want = %v", todos, []domains.Todo{})
  }

  todo, err := repo.Create(&domains.Todo{ Text: "明日から本気出す" }, time.Date(2021, 4, 1, 10, 0, 0, 0, time.UTC))
  if err != nil {
    t.Fatalf("setup failed by create. err = %v", err)
  }

  type args struct {
    todo *domains.Todo
    now time.Time
  }
  type want struct {
    id domains.TodoID
    text string
    createdAt time.Time
    updatedAt time.Time
  }
  tests := []struct {
    name string
    args args
    want want
  }{
    { name: "ok", args: args{ todo: &domains.Todo{ Id: todo.Id, Text: "今日から本気出す"}, now: time.Date(2021, 4, 2, 15, 0, 0, 0, time.UTC) }, want: want{ id: todo.Id, text: "今日から本気出す", createdAt: time.Date(2021, 4, 1, 10, 0, 0, 0, time.UTC), updatedAt: time.Date(2021, 4, 2, 15, 0, 0, 0, time.UTC) } },
  }

 for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      todo, err := repo.Update(tt.args.todo, tt.args.now)
      if err != nil {
        t.Fatalf("err = %v", err)
      }
      if todo.Id != tt.want.id {
        t.Fatalf("todo.Id = %vm want = %v", todo.Id, tt.want.id)
      }
      if todo.Text != tt.want.text {
        t.Fatalf("todo.Text = %v, want = %v", todo.Text, tt.want.text)
      }
      if todo.CreatedAt != tt.want.createdAt {
        t.Fatalf("todo.CreatedAt = %v, want = %v", todo.CreatedAt, tt.want.createdAt)
      }
      if todo.UpdatedAt != tt.want.updatedAt {
        t.Fatalf("todo.UpdatedAt = %v, want = %v", todo.UpdatedAt, tt.want.updatedAt)
      }
    })
  }
}

func TestTodoDelete (t *testing.T) {
  repo := domains.New()
  todo, err := repo.Create(&domains.Todo{ Text: "明日から本気出す" }, time.Date(2021, 4, 1, 10, 0, 0, 0, time.UTC))
  if err != nil {
    t.Fatalf("setup failed by create. err = %v", err)
  }

  type args struct {
    todo *domains.Todo
    now time.Time
  }
  type want struct {
    id domains.TodoID
    text string
  }
  tests := []struct {
    name string
    args args
    want want
  }{
    { name: "ok", args: args{ todo: &domains.Todo{ Id: todo.Id } }, want: want{ id: todo.Id, text: "今日から本気出す" } },
  }

 for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      todo, err := repo.Delete(tt.args.todo)
      if err != nil {
        t.Fatalf("err = %v", err)
      }
      if todo.Id != tt.want.id {
        t.Fatalf("todo.Id = %vm want = %v", todo.Id, tt.want.id)
      }
      todo, err = repo.Get(todo.Id)
      if err == nil {
        t.Fatalf("削除したTodoが取得できてしまう Todo = %v", todo)
      }
    })
  }
}
