package memberships

import (
	"context"
	"database/sql"
	"errors"
	"github.com/novinbukannopin/fc-simple-forum/internal/model/memberships"
)

func (r *Repository) GetUser(ctx context.Context, email, username string, userID int64) (*memberships.UserModel, error) {
	query := `SELECT id, email, password, username, created_at, updated_at, created_by, updated_by
	FROM fastcampus.users WHERE email = ? OR username = ? OR id = ?`
	rows := r.db.QueryRowContext(ctx, query, email, username, userID)

	var response memberships.UserModel
	err := rows.Scan(&response.ID, &response.Email, &response.Password, &response.Username, &response.CreatedAt, &response.UpdatedAt, &response.CreatedBy, &response.UpdatedBy)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &response, nil
}

func (r *Repository) CreateUser(ctx context.Context, user memberships.UserModel) error {
	query := `INSERT INTO fastcampus.users (email, username, password, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, user.Email, user.Username, user.Password, user.CreatedAt, user.UpdatedAt, user.CreatedBy, user.UpdatedBy)

	if err != nil {
		return err
	}
	return nil
}
