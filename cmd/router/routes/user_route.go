package routes

import (
	"net/http"

	"github.com/scalent-sushil/user-management-go/pkg/handler"
)

var userRoutes = []Route{

	Route{
		Url:          "/user",
		Method:       http.MethodGet,
		Handler:      handler.UserProfile,
		AuthRequired: true,
	},
	Route{
		Url:          "/user/update",
		Method:       http.MethodPut,
		Handler:      handler.UpdateUser,
		AuthRequired: true,
	},

	Route{
		Url:          "/user/profilepic",
		Method:       http.MethodPut,
		Handler:      handler.UploadImage,
		AuthRequired: true,
	},
	Route{
		Url:          "/user/changepassword",
		Method:       http.MethodPut,
		Handler:      handler.ChangePassword,
		AuthRequired: true,
	},
}
