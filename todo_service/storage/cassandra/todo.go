package cassandra

import (
	"fmt"
	"time"

	"github.com/abdukhashimov/blogcas/todo_service/storage/repo"
	"github.com/abdukhashimov/blogcas/todo_service/todopb"
	"github.com/gocql/gocql"
)

type todoRepo struct {
	session *gocql.Session
}

// it returns the sms Repo with session
func NewTodoRepo(session *gocql.Session) repo.TodoStorateI {
	return &todoRepo{session: session}
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

func (repo *todoRepo) Update(todo *todopb.Todo) (*todopb.Todo, error) {
	o, err := repo.Get(fmt.Sprint(todo.GetId()))
	if err != nil {
		return nil, err
	}
	query := repo.session.Query(
		`UPDATE todos SET 
			title, 
			description,
			done, 
			updatedAt 
			WHERE id = ?`,
		todo.GetTitle(),
		todo.GetDescription(),
		todo.GetDone(),
		time.Now(),
		o.GetId(),
	)

	if err := query.Exec(); err != nil {
		return nil, err
	}
	return todo, nil
}

func (repo *todoRepo) Delete(id string) error {
	o, err := repo.Get(id)

	if err != nil {
		return err
	}

	query := repo.session.Query(`DELETE from todos WHERE id = ?`, o.GetId())

	if err := query.Exec(); err != nil {
		return err
	}

	return nil
}

// func (repo *todoRepo) GetAll(*todopb.Empty) (*todopb.Todos, error) {
// 	query := repo.session.Query(`SELECT id, title, description, done, createdAt, updatedAt from todos`)
// 	defer query.Release()

// 	if err := query.Exec(); err != nil {
// 		return nil, err
// 	}
// 	return nil, nil
// }
