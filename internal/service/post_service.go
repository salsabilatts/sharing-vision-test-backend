package service

import (
	"errors"
	"strings"

	"github.com/salsabilatts/sharing-vision-test/internal/model"
	"github.com/salsabilatts/sharing-vision-test/internal/repository"
)

var (
	ErrTitleRequired    = errors.New("title is required and must be at least 20 characters")
	ErrContentRequired  = errors.New("content is required and must be at least 200 characters")
	ErrCategoryRequired = errors.New("category is required and must be at least 3 characters")
	ErrInvalidStatus    = errors.New("status must be publish, draft, or thrash")
	ErrPostNotFound     = errors.New("article not found")
)

type PostService struct {
	repository *repository.PostRepository
}

func NewPostService(repository *repository.PostRepository) *PostService {
	return &PostService{
		repository: repository,
	}
}

func (s *PostService) Create(post *model.Post) error {
	if err := validatePost(post); err != nil {
		return err
	}

	return s.repository.Create(post)
}

func (s *PostService) GetAll(limit, offset int, status string) ([]model.Post, error) {
	return s.repository.FindAll(limit, offset, status)
}

func (s *PostService) GetByID(id uint) (*model.Post, error) {
	post, err := s.repository.FindByID(id)

	if err != nil {
		return nil, ErrPostNotFound
	}

	return post, nil
}

func (s *PostService) Update(id uint, post *model.Post) (*model.Post, error) {
	existingPost, err := s.repository.FindByID(id)
	if err != nil {
		return nil, ErrPostNotFound
	}

	if err := validatePost(post); err != nil {
		return nil, err
	}

	existingPost.Title = post.Title
	existingPost.Content = post.Content
	existingPost.Category = post.Category
	existingPost.Status = post.Status

	if err := s.repository.Update(existingPost); err != nil {
		return nil, err
	}

	return existingPost, nil
}

func (s *PostService) Delete(id uint) error {
	post, err := s.repository.FindByID(id)

	if err != nil {
		return ErrPostNotFound
	}

	return s.repository.Delete(post)
}

func validatePost(post *model.Post) error {
	if strings.TrimSpace(post.Title) == "" || len([]rune(strings.TrimSpace(post.Title))) < 20 {
		return ErrTitleRequired
	}

	if strings.TrimSpace(post.Content) == "" || len([]rune(strings.TrimSpace(post.Content))) < 200 {
		return ErrContentRequired
	}

	if strings.TrimSpace(post.Category) == "" || len([]rune(strings.TrimSpace(post.Category))) < 3 {
		return ErrCategoryRequired
	}

	switch strings.ToLower(strings.TrimSpace(post.Status)) {
	case "publish", "draft", "thrash":
		post.Status = strings.ToLower(strings.TrimSpace(post.Status))
	default:
		return ErrInvalidStatus
	}

	return nil
}
