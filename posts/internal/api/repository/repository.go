package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/hunaisashraf/go-auth/internal/api/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	AddUser(name, email, phone, password string) error
	GetUser(email string) (model.User, error)
	EditUser(name, email, phone, password string) (model.User, error)
	AddPost(post model.Posts) error
	GetPost(id string) (model.Posts, error)
}

type Repo struct {
	conn *pgxpool.Pool
}

func NewRepository(conn *pgxpool.Pool) *Repo {
	return &Repo{conn: conn}
}

func (r *Repo) AddUser(name, email, phone, password string) error {
	query := `
	INSERT INTO TABLE users(
	name,email,phone,password)
	values($1,$2,$3,$4);`
	_, err := r.conn.Exec(context.Background(), query, name, email, phone, password)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (r *Repo) GetUser(email string) (model.User, error) {
	query := `SELECT * FROM users WHERE email=$1;`

	row := r.conn.QueryRow(context.Background(), query, email)
	fmt.Println(row)

	return model.User{}, nil
}

func (r *Repo) EditUser(name, email, phone, password string) (model.User, error) {
	query := "UPDATE users WHERE email=$1 SET name=$2, email=$3,phone=$3,password=$4"

	_, err := r.conn.Exec(context.Background(), query, name, email, phone, password)
	if err != nil {
		log.Fatal(err)
	}
	return model.User{}, nil
}

func (r *Repo) AddPost(post model.Posts) error {
	query := `INSERT INTO posts(title,body,userid) values($1,$2,$3)`

	_, err := r.conn.Exec(context.Background(), query, post.Title, post.Body, post.UserId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (r *Repo) GetPost(id string) (model.Posts, error) {
	query := `SELECT * FROM posts WHERE id=$1`
	fmt.Println("id is ", id)
	var post model.Posts

	err := r.conn.QueryRow(context.Background(), query, id).Scan(&post.Id, &post.Title, &post.Body, &post.UserId)

	if err != nil {
		fmt.Println(err)
		return model.Posts{}, err
	}

	return post, nil
}
