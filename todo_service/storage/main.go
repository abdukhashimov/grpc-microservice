package storage

import (
	"github.com/abdukhashimov/blogcas/todo_service/storage/cassandra"
	"github.com/abdukhashimov/blogcas/todo_service/storage/repo"
	"github.com/gocql/gocql"
)

// StorateI ...
type StorageI interface {
	Todo() repo.TodoStorateI
}

type storageCassandra struct {
	session  *gocql.Session
	todoRepo repo.TodoStorateI
}

func newStorageCassandra(session *gocql.Session) StorageI {
	return &storageCassandra{session: session, todoRepo: cassandra.}
}
