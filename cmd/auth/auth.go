package auth

import (
	"github.com/scalent-sushil/user-management-go/cmd/security"
	"github.com/scalent-sushil/user-management-go/database"
	"github.com/scalent-sushil/user-management-go/pkg/models"
)

// SignIn fuction is use to signIn only if user is active
func SignIn(email, password string) (string, error) {

	user := models.User{}
	var err error
	// var db *gorm.DB
	// db, err = database.Connect()
	// if err != nil {
	// 	ch <- false
	// 	return
	// }
	err = database.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	if user.Status == "Deactivated" {
		return "", err
	}
	pass := security.VerifyPassword(user.Password, password)
	if pass == false {
		return "", err
	}

	return CreateToken(uint32(user.ID), user.UserType)
}

//SignUp function is use to create new user when user is registering himself
// and it return the token of the user
func SignUp(email, password string) (string, error) {

	user := models.User{}
	user.Email = email
	user.Password = password
	user.Status = "Activated"
	var err error

	err = database.DB.Debug().Model(&models.User{}).Create(&user).Error
	if err != nil {
		return "", err
	}

	return CreateToken(uint32(user.ID), user.UserType)
}

//AdminSignIn is signIn Api for admin
func AdminSignIn(email, password string) (string, error) {

	user := models.User{}
	var err error
	// var db *gorm.DB
	err = database.DB.Debug().Model(models.User{}).Where("email = ? AND user_type = ? ", email, "admin").Take(&user).Error
	if err != nil {
		return "", err
	}
	pass := security.VerifyPassword(user.Password, password)
	if pass == false {
		return "", err
	}

	return CreateToken(uint32(user.ID), user.UserType)
}

// EmailPassword to recover/new password
func EmailPassword(email string) bool {

	// This fuction checks the user email at the time of forgot password api that user exist or not

	user := models.User{}
	var err error
	// var db *gorm.DB
	err = database.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return false
	}
	return true
}

// SetPassword is fuction call after email password
func SetPassword(email, password string) bool {

	//this function is use set new password after user is sent opt on its email

	var err error
	hashedPassword, _ := security.Hash(password)
	err = database.DB.Debug().Model(models.User{}).Where("email = ?", email).UpdateColumns(
		map[string]interface{}{
			"password": hashedPassword,
		},
	).Error
	if err != nil {
		return false
	}

	return true
}
