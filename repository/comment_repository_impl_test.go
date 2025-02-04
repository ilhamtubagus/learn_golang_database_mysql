package repository

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // import mysql driver to prevent errors unknown driver
	"github.com/ilhamtubagus/learn_golang_database_mysql"
	"github.com/ilhamtubagus/learn_golang_database_mysql/entity"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	db, _ := learn_golang_database_mysql.GetConnection()
	commentRepository := NewCommentRepository(db)

	ctx := context.Background()
	comment := entity.Comment{Email: "test@example.com", Comment: "This is a test comment"}

	result, err := commentRepository.InsertComment(ctx, &comment)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(result)
}

func TestCommentFindById(t *testing.T) {
	db, _ := learn_golang_database_mysql.GetConnection()
	commentRepository := NewCommentRepository(db)

	ctx := context.Background()

	result, err := commentRepository.FindCommentById(ctx, 100)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(result)
}

func TestCommentFindAll(t *testing.T) {
	db, _ := learn_golang_database_mysql.GetConnection()
	commentRepository := NewCommentRepository(db)

	ctx := context.Background()

	comments, err := commentRepository.FindAllComments(ctx)
	if err != nil {
		t.Fatal(err)
	}

	for _, comment := range *comments {
		fmt.Println(comment)
	}
}
