package domains

import (
  "time"
  "errors"
  "github.com/google/uuid"
)

type TodoID string

type Todo struct {
  Id TodoID
  Text string
  CreatedAt time.Time
  UpdatedAt time.Time
}

type TodoRepository interface {
  Get(id TodoID) (*Todo, error)
  GetAll() ([]*Todo, error)
  Create(*Todo, time.Time) (*Todo, error)
  Update(*Todo, time.Time) (*Todo, error)
  Delete(*Todo) (*Todo, error)
}

func New() TodoRepository {
  return &todoRepository {
    database: map[TodoID]*Todo{},
  }
}

type todoRepository struct {
  database map[TodoID]*Todo
}

func (repo *todoRepository) Create (todo *Todo, now time.Time) (*Todo, error) {
  todo.Id = TodoID(uuid.New().String())
  todo.CreatedAt = now

  repo.database[todo.Id] = todo
  return todo, nil
}

func (repo *todoRepository) Delete (todo *Todo) (*Todo, error) {
  if todo.Id == "" {
    return nil, errors.New("idが指定されてない")
  }

  delete(repo.database, todo.Id)

  return todo, nil
}

func (repo *todoRepository) Get (id TodoID) (*Todo, error) {
  if id == "" {
    return nil, errors.New("idが指定されていない")
  }

  todo, ok := repo.database[id]
  if !ok {
    return nil, errors.New("Todoが見つからない")
  }
  return todo, nil
}

func (repo *todoRepository) GetAll () ([]*Todo, error) {
  todos := make([]*Todo, 0, len(repo.database))
  for _, todo := range repo.database {
    todos = append(todos, todo)
  }

  return todos, nil
}

func (repo *todoRepository) Update (todo *Todo, now time.Time) (*Todo, error) {
  if todo.Id == "" {
    return nil, errors.New("idが指定されてない")
  }

  old, ok := repo.database[todo.Id]
  if !ok {
    return nil, errors.New("Todoが見つからない")
  }

  if old.Text != todo.Text {
    old.Text = todo.Text
    old.UpdatedAt = now
  }

  return old, nil
}

