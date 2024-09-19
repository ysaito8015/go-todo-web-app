package store

import (
	"errors"

	"github.com/ysaito8015/go-todo-api/entity"
)

var (
	Tasks = &TaskStore{Tasks: map[int]*entity.Task{}}

	ErrNotFound = errors.New("not found")
)

type TaskStore struct {
	// export the fields for fake implementation
	LastID entity.TaskID
	Tasks  map[entity.TaskID]*entity.Task
}

func (s *TaskStore) Add(t *entity.Task) (int, error) {
	ts.LastID++
	t.ID = ts.LastID
	ts.Tasks[t.ID] = t
	return t.ID, nil
}

// All methods return all tasks
func (ts *TaskStore) All() entity.Tasks {
	tasks := make([]*entity.Task, len(t.Tasks))
	for i, t := range ts.Tasks {
		tasks[i-1] = t
	}
	return tasks
}
