package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/yadavsushil07/user-management/api/auth"
	"github.com/yadavsushil07/user-management/api/database"
	"github.com/yadavsushil07/user-management/api/models"
	"github.com/yadavsushil07/user-management/api/repository"
	"github.com/yadavsushil07/user-management/api/repository/crud"
	"github.com/yadavsushil07/user-management/api/responses"

	"github.com/gorilla/mux"
)

// CreateBlog it creates new blog...
func CreateBlog(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	blog := models.Blog{}
	err = json.Unmarshal(body, &blog)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userID, _, err := auth.ExtractClaim(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	blog.AuthorID = int(userID)
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryBlogsCURD(db)

	func(blogRepository repository.BlogRepository) {
		blog, err := blogRepository.Save(blog)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		// w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, user.ID))
		responses.JSON(w, http.StatusCreated, blog)
	}(repo)
}

// GetBlogs , all blog of user
func GetBlogs(w http.ResponseWriter, r *http.Request) {

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryBlogsCURD(db)

	func(blogRepository repository.BlogRepository) {
		blogs, err := blogRepository.FindAll()
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		responses.JSON(w, http.StatusOK, blogs)
	}(repo)
}

// GetBlog blog by id...
func GetBlog(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryBlogsCURD(db)

	func(blogRepository repository.BlogRepository) {
		blog, err := blogRepository.FindByID(uint32(bid))
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		responses.JSON(w, http.StatusOK, blog)
	}(repo)
}

// Updateblog updates the blog
func Updateblog(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	blog := models.Blog{}
	err = json.Unmarshal(body, &blog)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userID, _, err := auth.ExtractClaim(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	blog.AuthorID = int(userID)
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryBlogsCURD(db)

	func(blogRepository repository.BlogRepository) {
		status, err := blogRepository.UpdateBlog(uint32(bid), blog)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		responses.JSON(w, http.StatusOK, status)
	}(repo)
}

// Deleteblog delete's the blog of the user...
func Deleteblog(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryBlogsCURD(db)

	func(blogRepository repository.BlogRepository) {
		str, err := blogRepository.DeleteBlog(uint32(bid))
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		responses.JSON(w, http.StatusNoContent, str)
	}(repo)
}
