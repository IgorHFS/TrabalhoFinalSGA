package services

import "github.com/jackc/pgx/v5/pgconn"

type Error struct {
	Status  int
	Message string
}

func (e *Error) Error() string     { return e.Message }
func BadRequest(m string) error    { return &Error{400, m} }
func NotFound(m string) error      { return &Error{404, m} }
func Conflict(m string) error      { return &Error{409, m} }
func Unprocessable(m string) error { return &Error{422, m} }
func isUnique(err error) bool      { p, ok := err.(*pgconn.PgError); return ok && p.Code == "23505" }
