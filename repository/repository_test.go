package repository

import (
	"context"
	"fmt"
	gomysql "go-mysql"
	"go-mysql/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T){
	commentRepository := NewCommentRepository(gomysql.GetConnection())
	ctx := context.Background()
	comment := entity.Comment{
		Email: "repository@test.com",
		Comment: "Test Repository",
	}

	result, err := commentRepository.Insert(ctx, comment)
	
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

}

func TestFindById(t *testing.T){
	cmtRepo := NewCommentRepository(gomysql.GetConnection())

	ctx := context.Background()

	comment, err := cmtRepo.FindById(ctx, 110022)

	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T){
	cmtRepo := NewCommentRepository(gomysql.GetConnection())

	ctx := context.Background()

	comment, err := cmtRepo.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}