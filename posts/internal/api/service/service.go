package services

import (
	"fmt"

	"github.com/hunaisashraf/go-auth/internal/api/model"
	"github.com/hunaisashraf/go-auth/internal/api/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

func (r *Service) CreatePost(post model.Posts) (model.Posts, error) {

	err := r.repo.AddPost(post)
	if err != nil {
		fmt.Println(err)
		return model.Posts{}, err
	}

	return post, nil
}

func (r *Service) GetPost(postId string) (model.Posts, error) {
	fmt.Println(postId)

	post, err := r.repo.GetPost(postId)

	if err != nil {
		return model.Posts{}, err
	}

	return post, nil
}

// func (r *Service) Signup(user model.User) {

// 	r.repo.GetUser(user.Email)

// 	err := r.repo.AddUser(user.Username, user.Email, user.Phone, user.Password)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }
