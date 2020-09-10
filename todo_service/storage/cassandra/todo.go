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
		createdAt,
		updatedAt,
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

func (repo *todoRepo) Get(id string) (*todopb.Todo, error) {
	var todo todopb.Todo
	uuid, err := gocql.ParseUUID(id)

	if err != nil {
		return nil, err
	}

	query := repo.session.Query(`SELECT 
		id,
		title,
		description,
		done,
		createdAt,
		updatedAt FROM todos WHERE id = ?`, uuid)
	var createdAt time.Time
	var updatedAt time.Time

	if err := query.Scan(
		&todo.Id,
		&todo.Title,
		&todo.Description,
		&createdAt,
		&updatedAt,
	); err != nil {
		return nil, err
	}
	todo.CreatedAt = createdAt.String()
	todo.UpdatedAt = updatedAt.String()

	return &todo, nil
}
