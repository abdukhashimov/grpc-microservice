package storage

import (
	"github.com/abdukhashimov/blogcas/blog_service/storage/repo"
)

type StorageI interface {
	Todo() repo.TodoStorateI
}
