package repository

import "github.com/scalent-sushil/user-management-go/pkg/models"

type BlogRepository interface {
	Save(models.Blog) (models.Blog, error)
	FindAll() ([]models.Blog, error)
	FindByID(uint32) (models.Blog, error)
	UpdateBlog(uint32, models.Blog) (string, error)
	DeleteBlog(uint32) (string, error)
}
