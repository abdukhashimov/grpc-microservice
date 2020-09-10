package cassandra

import (
	"github.com/gocql/gocql"
)

type todoRepo struct {
	session *gocql.Session
}
