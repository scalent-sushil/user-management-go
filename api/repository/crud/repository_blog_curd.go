package crud

import (
	"errors"
	"time"

	"github.com/yadavsushil07/user-management/api/models"
	"github.com/yadavsushil07/user-management/api/utils/channels"

	"gorm.io/gorm"
)

type repositoryBlogsCRUD struct {
	db *gorm.DB
}

// NewRepositoryBlogsCURD this is use to return the database after the curd method
func NewRepositoryBlogsCURD(db *gorm.DB) *repositoryBlogsCRUD {
	return &repositoryBlogsCRUD{db}
}

// Save it is use to create new blog
func (r *repositoryBlogsCRUD) Save(blog models.Blog) (models.Blog, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Blog{}).Create(&blog).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return blog, nil
	}
	return models.Blog{}, err
}

// FindAll this is use to get all blogs of the user
func (r *repositoryBlogsCRUD) FindAll() ([]models.Blog, error) {
	var err error
	blogs := []models.Blog{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Blog{}).Limit(10).Find(&blogs).Error
		if err != nil {
			ch <- false
			return
		}
		if len(blogs) > 0 {
			for i := range blogs {
				err = r.db.Debug().Model(&models.Blog{}).Where("id = ?", blogs[i].AuthorID).Error
				if err != nil {
					ch <- false
					return
				}
			}
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return blogs, nil
	}
	return nil, err
}

// FindByID this function is use to find blog by id
func (r *repositoryBlogsCRUD) FindByID(bid uint32) (models.Blog, error) {
	var err error
	blogs := models.Blog{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Blog{}).Where("id = ?", bid).Take(&blogs).Error
		if err != nil {
			ch <- false
			return
		}
		if blogs.ID != 0 {
			err = r.db.Debug().Model(&models.Blog{}).Where("id = ?", blogs.AuthorID).Error
			if err != nil {
				ch <- false
				return
			}
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return blogs, nil
	}
	return models.Blog{}, err
}

// UpdateBlog this function is use to update the title and the content of the blog
func (r *repositoryBlogsCRUD) UpdateBlog(bid uint32, blogs models.Blog) (string, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.Blog{}).Where("id = ?", bid).Take(&models.Blog{}).UpdateColumns(
			map[string]interface{}{
				"title":      blogs.Title,
				"content":    blogs.Content,
				"updated_at": time.Now(),
			},
		)

		ch <- true
	}(done)
	if channels.OK(done) {
		if rs.Error != nil {
			return "", errors.New("record not found")
		}
		return "successfull update blog", nil
	}
	return "", rs.Error
}

// DeleteBlog this function is use to delete blog
func (r *repositoryBlogsCRUD) DeleteBlog(bid uint32) (string, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.Blog{}).Where("id = ?", bid).Take(&models.Blog{}).Delete(&models.Blog{})
		if rs.Error != nil {
			ch <- false
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return "successfull update blog", nil
	}
	return "Record not found", rs.Error
}
