package repo

import (
	"github.com/abdukhashimov/blogcas/blog_service/todopb"
)

tpye TodoStorateI interface {
    Create(todo *todopb.Todo) (string, error)
    Get(id string) (*todopb.Todo, error)
    Update(todo *todopb.Todo) (*todopb.Todo, error)
    Delete(id string) error
    GetAll(todopb.Empty) (*todopb.Todos, error)
}