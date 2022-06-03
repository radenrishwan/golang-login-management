package repository

import (
	"context"
	"database/sql"
	"errors"
	"login-management/helper"
	"login-management/model/domain"
)

type SessionRepository interface {
	Save(ctx context.Context, tx *sql.Tx, session domain.Session) domain.Session
	FindById(ctx context.Context, tx *sql.Tx, session domain.Session) (domain.Session, error)
	Delete(ctx context.Context, tx *sql.Tx, session domain.Session) domain.Session
}

type sessionRepository struct {
}

func NewSessionRepository() SessionRepository {
	return &sessionRepository{}
}

func (repository *sessionRepository) Save(ctx context.Context, tx *sql.Tx, session domain.Session) domain.Session {
	query := "insert into sessions values (?, ?, ?, ?)"

	_, err := tx.ExecContext(ctx, query, session.Id, session.UserId, session.CreatedAt, session.ExpiredAt)
	helper.PanicIfError(err)

	return session
}

func (repository *sessionRepository) FindById(ctx context.Context, tx *sql.Tx, session domain.Session) (domain.Session, error) {
	query := " select * from sessions where user_id = ?"

	rows, err := tx.QueryContext(ctx, query, session.UserId)
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(session.Id, session.UserId, session.CreatedAt, session.ExpiredAt)
		helper.PanicIfError(err)

		return session, nil
	}

	return session, errors.New("user not found")
}

func (repository *sessionRepository) Delete(ctx context.Context, tx *sql.Tx, session domain.Session) domain.Session {
	query := "delete from sessions where user_id = ?"

	_, err := tx.ExecContext(ctx, query, session.UserId)
	helper.PanicIfError(err)

	return session
}
