package posts

import (
	"context"
	"github.com/novinbukannopin/fc-simple-forum/internal/model/posts"
)

func (r *Repository) CreateComment(ctx context.Context, model posts.CommentModel) error {
	query := `INSERT INTO fastcampus.comments (post_id, user_id, comment_content, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, model.PostId, model.UserId, model.CommentContent, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetCommentByPostId(ctx context.Context, postId int64) ([]posts.Comment, error) {
	query := `SELECT c.id, c.user_id, c.comment_content, u.username FROM fastcampus.comments c JOIN fastcampus.users u ON c.user_id = u.id WHERE c.post_id = ?`

	rows, err := r.db.QueryContext(ctx, query, postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	response := make([]posts.Comment, 0)
	for rows.Next() {
		var (
			comment  posts.Comment
			username string
		)
		err = rows.Scan(&comment.Id, &comment.UserId, &comment.CommentContent, &username)
		if err != nil {
			return nil, err
		}
		response = append(response, posts.Comment{
			Id:             comment.Id,
			UserId:         comment.UserId,
			CommentContent: comment.CommentContent,
			Username:       username,
		})
	}
	return response, nil
}
