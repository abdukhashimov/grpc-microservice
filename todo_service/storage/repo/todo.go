package repo

import "github.com/abdukhashimov/blogcas/todo_service/todopb"

/*
 *	Todo Storage interface includes the functions that must be implemented in
 *	Storage that is inheriting it.
 */
type TodoStorateI interface {
	Create(todo *todopb.Todo) (string, error)
	Get(id string) (*todopb.Todo, error)
	Update(todo *todopb.Todo) (*todopb.Todo, error)
	Delete(id string) error
	GetAll(todopb.Empty) (*todopb.Todos, error)
}
