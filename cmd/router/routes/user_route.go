package routes

import (
	"net/http"

	controllers "github.com/scalent-sushil/user-management-go/pkg/handler"
)

var userRoutes = []Route{

	Route{
		Url:          "/user",
		Method:       http.MethodGet,
		Handler:      controllers.UserProfile,
		AuthRequired: true,
	},
	Route{
		Url:          "/user/update",
		Method:       http.MethodPut,
		Handler:      controllers.UpdateUser,
		AuthRequired: true,
	},

	Route{
		Url:          "/user/profilepic",
		Method:       http.MethodPut,
		Handler:      controllers.UploadImage,
		AuthRequired: true,
	},
	Route{
		Url:          "/user/changepassword",
		Method:       http.MethodPut,
		Handler:      controllers.ChangePassword,
		AuthRequired: true,
	},
}
