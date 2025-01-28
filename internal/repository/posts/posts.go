package posts

import (
	"context"
	"github.com/novinbukannopin/fc-simple-forum/internal/model/posts"
)

func (r *Repository) CreatePost(ctx context.Context, model posts.PostModel) error {
	query := `INSERT INTO fastcampus.posts (user_id, post_title, post_content, post_hashtags, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, model.UserId, model.PostTitle, model.PostContent, model.PostHashtags, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)

	if err != nil {
		return err
	}

	return nil
}
