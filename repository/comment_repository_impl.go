package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-mysql/entity"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repo *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	result, err := repo.DB.ExecContext(ctx, script, comment.Email, comment.Comment)
	// fmt.Println("insert", err)
	defer repo.DB.Close()
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

func (repo *commentRepositoryImpl) FindById (ctx context.Context, id int32) (entity.Comment, error){
	script := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repo.DB.QueryContext(ctx, script, id)
	defer repo.DB.Close()
	defer rows.Close()

	comment := entity.Comment{}
	
	if err != nil {
		return comment, err
	}
	
	if rows.Next(){
		// exist
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		// not found
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}

func (repo *commentRepositoryImpl) FindAll (ctx context.Context) ([]entity.Comment, error){
	script := "SELECT id, email, comment FROM comments"
	rows, err := repo.DB.QueryContext(ctx, script)
	defer repo.DB.Close()
	defer rows.Close()

	if err != nil {
		return nil, err
	}
	
	var comments []entity.Comment
	for rows.Next(){
		// exist
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)

		comments = append(comments, comment)
	}
	return comments, nil
}