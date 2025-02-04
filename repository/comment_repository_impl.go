package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ilhamtubagus/learn_golang_database_mysql/entity"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) InsertComment(ctx context.Context, comment *entity.Comment) (*entity.Comment, error) {
	script := "INSERT INTO comments(email, comment) VALUES (?,?)"
	result, err := repository.DB.ExecContext(ctx, script, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}

	comment.Id = int32(id)

	return comment, nil
}

func (repository *commentRepositoryImpl) FindCommentById(ctx context.Context, id int32) (*entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments WHERE id =? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	comment := &entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		if err != nil {
			return nil, err
		}
		return comment, nil
	} else {
		return comment, errors.New("ID not found " + strconv.Itoa(int(id)))
	}
}

func (repository *commentRepositoryImpl) FindAllComments(ctx context.Context) (*[]entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		err := rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return &comments, nil
}
