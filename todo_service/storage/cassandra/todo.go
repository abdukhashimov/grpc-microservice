package cassandra

import (
	"time"

	"github.com/abdukhashimov/blogcas/todo_service/todopb"
	"github.com/gocql/gocql"
)

type todoRepo struct {
	session *gocql.Session
}

func (repo *todoRepo) Create(todo *todopb.Todo) (string, error) {
	id, err := gocql.RandomUUID()
	if err != nil {
		return "", nil
	}

	query := repo.session.Query(`INSERT INTO todos (
		id,
		title,
		description,
		done,
		created_at,
		updated_at,
	) VALUES (?, ?, ?, ?, ?, ?)`,
		id,
		todo.GetTitle(),
		todo.GetDescription(),
		time.Now(),
		time.Now(),
	)
	if err := query.Exec(); err != nil {
		return "", nil
	}
	return id.String(), nil
}
