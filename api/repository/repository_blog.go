package repository

import "github.com/yadavsushil07/user-management/api/models"

type BlogRepository interface {
	Save(models.Blog) (models.Blog, error)
	FindAll() ([]models.Blog, error)
	FindByID(uint32) (models.Blog, error)
	UpdateBlog(uint32, models.Blog) (string, error)
	DeleteBlog(uint32) (string, error)
}
