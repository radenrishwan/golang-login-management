package repository

import (
	"context"
	"database/sql"
	"errors"
	"login-management/helper"
	"login-management/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	FindById(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	FindByUsername(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (repository *userRepository) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	query := "insert into users values (?, ?, ?, ?, ?, ?, ?)"

	_, err := tx.ExecContext(ctx, query, user.Id, user.Email, user.Username, user.Password, user.Role, user.CreatedAt, user.UpdatedAt)
	helper.PanicIfError(err)

	return user
}

func (repository *userRepository) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	query := "update users set email = ?, username = ?, password = ?, updated_at = ? where id = ?"

	_, err := tx.ExecContext(ctx, query, user.Email, user.Username, user.Password, user.UpdatedAt, user.Id)
	helper.PanicIfError(err)

	return user
}

func (repository *userRepository) Delete(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	query := "delete from users where id = ?"

	_, err := tx.ExecContext(ctx, query, user.Id)
	helper.PanicIfError(err)

	return user
}

func (repository *userRepository) FindById(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	query := " select * from users where id = ?"

	rows, err := tx.QueryContext(ctx, query, user.Id)
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		helper.PanicIfError(err)

		return user, nil
	}

	return user, errors.New("user not found")
}

func (repository *userRepository) FindByUsername(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	query := " select * from users where username = ?"

	rows, err := tx.QueryContext(ctx, query, user.Username)
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		helper.PanicIfError(err)

		return user, nil
	}

	return user, errors.New("user not found")
}
