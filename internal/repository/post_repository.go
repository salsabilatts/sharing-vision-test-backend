package repository

import (
	"gorm.io/gorm"

	"github.com/salsabilatts/sharing-vision-test/internal/model"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (r *PostRepository) Create(post *model.Post) error {
	return r.db.Create(post).Error
}

func (r *PostRepository) FindAll(limit, offset int, status string) ([]model.Post, error) {
	var posts []model.Post

	query := r.db

	if status != "" {
		query = query.Where("status = ?", status)
	}

	err := query.
		Limit(limit).
		Offset(offset).
		Order("created_date DESC").
		Find(&posts).Error

	return posts, err
}

func (r *PostRepository) FindByID(id uint) (*model.Post, error) {
	var post model.Post

	err := r.db.First(&post, id).Error
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *PostRepository) Update(post *model.Post) error {
	return r.db.Save(post).Error
}

func (r *PostRepository) Delete(post *model.Post) error {
	return r.db.Delete(post).Error
}
