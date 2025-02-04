package repository

import (
	"context"
	"github.com/ilhamtubagus/learn_golang_database_mysql/entity"
)

type CommentRepository interface {
	InsertComment(ctx context.Context, comment *entity.Comment) (*entity.Comment, error)
	FindCommentById(ctx context.Context, id int32) (*entity.Comment, error)
	FindAllComments(ctx context.Context) (*[]entity.Comment, error)
}
