package store

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

const (
	uniqueViolationCode = "23505"
)

func IsUniqueViolation(err error, constraint string) bool {
	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) {
		return false
	}

	if pgErr.Code != uniqueViolationCode {
		return false
	}

	if constraint == "" {
		return true
	}

	return pgErr.ConstraintName == constraint
}
