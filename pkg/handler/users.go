package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/scalent-sushil/user-management-go/cmd/auth"
	"github.com/scalent-sushil/user-management-go/cmd/responses"
	"github.com/scalent-sushil/user-management-go/database"
	"github.com/scalent-sushil/user-management-go/pkg/models"
	"github.com/scalent-sushil/user-management-go/pkg/repository"
	"github.com/scalent-sushil/user-management-go/pkg/repository/crud"

	"github.com/gorilla/mux"
)

// UserProfile function return the user profile by taking the userID as parameter ..
func UserProfile(w http.ResponseWriter, r *http.Request) {

	// db, err := database.Connect()
	// if err != nil {
	// 	responses.ERROR(w, http.StatusInternalServerError, err)
	// 	return
	// }

	// defer db.Close()
	userID, _, err := auth.ExtractClaim(r) // userID is fetch from the token of the user loged in
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	repo := crud.NewRepositoryUsersCURD(database.DB)
	func(UserRepository repository.UserRepository) {
		user, err := UserRepository.User(uint32(userID)) //userID pased as parameter
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		responses.JSON(w, http.StatusOK, user)
	}(repo)
}

//UploadImage function returns the filename(to save in database) of the saved file or an error if it occurs
func UploadImage(w http.ResponseWriter, r *http.Request) {

	name, err := FileUpload(r)
	if err != nil {
		responses.ERROR(w, http.StatusNoContent, err)
		return
	}
	user := models.User{}
	// db, err := database.Connect()
	// if err != nil {
	// 	responses.ERROR(w, http.StatusInternalServerError, err)
	// 	return
	// }
	userID, _, err := auth.ExtractClaim(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	repo := crud.NewRepositoryUsersCURD(database.DB)
	fmt.Println(user.ID)

	func(UserRepository repository.UserRepository) {
		user.ProfilePic = name
		user, err := UserRepository.UploadPic(uint32(userID), user)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusCreated, user)
	}(repo)
	fmt.Println(name)
	fmt.Println(err)

}

//FileUpload funtion ready the file as formfile and store the file where specify
func FileUpload(r *http.Request) (string, error) {

	r.ParseMultipartForm(32 << 10)

	file, handler, err := r.FormFile("file") //retrieve the file from form data

	fmt.Println(handler)
	if err != nil {
		return "", err
	}
	defer file.Close() //close the file when we finish

	f, err := os.OpenFile("upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer f.Close()
	io.Copy(f, file)

	return handler.Filename, nil
}

// ChangePassword function is use to update password when user is loged in
func ChangePassword(w http.ResponseWriter, r *http.Request) {

	// it checks the use by token

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	userID, _, err := auth.ExtractClaim(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	// db, err := database.Connect()
	// if err != nil {
	// 	responses.ERROR(w, http.StatusInternalServerError, err)
	// 	return
	// }

	// defer db.Close()

	repo := crud.NewRepositoryUsersCURD(database.DB)
	func(UserRepository repository.UserRepository) {
		rows, err := UserRepository.ResetPassword(userID, user)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		responses.JSON(w, http.StatusOK, rows)
	}(repo)

}

// UpdateUser function is use to update user details like name
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	// vars := mux.Vars(r)
	// id, err := strconv.ParseUint(vars["id"], 10, 32)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusBadRequest, err)
	// 	return
	// }

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	userID, _, err := auth.ExtractClaim(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	// db, err := database.Connect()
	// if err != nil {
	// 	responses.ERROR(w, http.StatusInternalServerError, err)
	// 	return
	// }

	repo := crud.NewRepositoryUsersCURD(database.DB)
	// if userID != uint32(id) {
	// 	responses.ERROR(w, http.StatusUnauthorized, err)
	// 	return
	// }

	func(UserRepository repository.UserRepository) {
		rows, err := UserRepository.Update(userID, user)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		responses.JSON(w, http.StatusOK, rows)
	}(repo)

}

//ADMIN PART

// AdminGetUsers function is use to by admin to fetch all users
func AdminGetUsers(w http.ResponseWriter, r *http.Request) {

	// db, err := database.Connect()
	// if err != nil {
	// 	responses.ERROR(w, http.StatusInternalServerError, err)
	// 	return
	// }

	_, userType, err := auth.ExtractClaim(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	repo := crud.NewRepositoryUsersCURD(database.DB)
	if userType == "admin" {
		func(UserRepository repository.UserRepository) {
			users, err := UserRepository.FindAll()
			if err != nil {
				responses.ERROR(w, http.StatusUnprocessableEntity, err)
				return
			}

			responses.JSON(w, http.StatusOK, users)
		}(repo)
	}
	if userType == "user" {
		responses.JSON(w, http.StatusUnauthorized, err)
		return
	}
}

// AdminUploadImage function is use to upload profile pic of admin
func AdminUploadImage(w http.ResponseWriter, r *http.Request) {

	//this function returns the filename(to save in database) of the saved file or an error if it occurs

	name, err := FileUpload(r)
	user := models.User{}
	// db, err := database.Connect()
	// if err != nil {
	// 	responses.ERROR(w, http.StatusInternalServerError, err)
	// 	return
	// }
	_, userType, err := auth.ExtractClaim(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	repo := crud.NewRepositoryUsersCURD(database.DB)
	if userType == "admin" {
		func(UserRepository repository.UserRepository) {
			user.ProfilePic = name
			user, err := UserRepository.Update(uint32(1), user)
			if err != nil {
				responses.ERROR(w, http.StatusUnprocessableEntity, err)
				return
			}
			// w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI))
			responses.JSON(w, http.StatusCreated, user)
		}(repo)
	}
	if userType == "user" {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
}

// AdminProfile function return the profile of the admin
func AdminProfile(w http.ResponseWriter, r *http.Request) {

	// db, err := database.Connect()
	// if err != nil {
	// 	responses.ERROR(w, http.StatusInternalServerError, err)
	// 	return
	// }

	// defer db.Close()
	_, userType, err := auth.ExtractClaim(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	repo := crud.NewRepositoryUsersCURD(database.DB)
	if userType == "admin" {
		func(UserRepository repository.UserRepository) {
			admin, err := UserRepository.Admin()
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			responses.JSON(w, http.StatusOK, admin)
		}(repo)
	}
	if userType == "user" {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
}

//GetUser fuction is use to fetch th profile of indiviual user
func GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	// db, err := database.Connect()
	// if err != nil {
	// 	responses.ERROR(w, http.StatusInternalServerError, err)
	// 	return
	// }
	_, userType, err := auth.ExtractClaim(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	repo := crud.NewRepositoryUsersCURD(database.DB)

	if userType == "admin" {
		func(UserRepository repository.UserRepository) {
			user, err := UserRepository.FindUserById(uint32(id))
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}
			responses.JSON(w, http.StatusOK, user)
		}(repo)
	}
}

// AdminUpdateUser function is to update detials of admin like name
func AdminUpdateUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// db, err := database.Connect()
	// if err != nil {
	// 	responses.ERROR(w, http.StatusInternalServerError, err)
	// 	return
	// }

	// defer db.Close()
	_, userType, err := auth.ExtractClaim(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	repo := crud.NewRepositoryUsersCURD(database.DB)
	if userType == "admin" {
		func(UserRepository repository.UserRepository) {
			rows, err := UserRepository.UpdateByAdmin(uint32(id), user)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			responses.JSON(w, http.StatusOK, rows)
		}(repo)
	}
	if userType == "user" {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
}

//DeleteUserByAdmin is use for delete/deactivate user by admin
func DeleteUserByAdmin(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// db, err := database.Connect()
	// if err != nil {
	// 	responses.ERROR(w, http.StatusInternalServerError, err)
	// 	return
	// }

	_, userType, err := auth.ExtractClaim(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	repo := crud.NewRepositoryUsersCURD(database.DB)
	if userType == "admin" {
		func(UserRepository repository.UserRepository) {
			rows, err := UserRepository.DeleteByAdmin(uint32(id), user)
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)
				return
			}

			responses.JSON(w, http.StatusOK, rows)
		}(repo)
	}
	if userType == "user" {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
}
