package routes

import (
	"net/http"

	controllers "github.com/scalent-sushil/user-management-go/pkg/handler"
)

var AdminRoutes = []Route{

	Route{
		Url:          "/admin/users",
		Method:       http.MethodGet,
		Handler:      controllers.AdminGetUsers,
		AuthRequired: true,
	},
	Route{
		Url:          "/admin",
		Method:       http.MethodGet,
		Handler:      controllers.AdminProfile,
		AuthRequired: true,
	},
	Route{
		Url:          "/admin/user/{id}",
		Method:       http.MethodGet,
		Handler:      controllers.GetUser,
		AuthRequired: true,
	},
	Route{
		Url:          "/admin/user/{id}",
		Method:       http.MethodPut,
		Handler:      controllers.AdminUpdateUser,
		AuthRequired: true,
	},

	Route{
		Url:          "/admin/user/{id}/status",
		Method:       http.MethodPut,
		Handler:      controllers.DeleteUserByAdmin,
		AuthRequired: true,
	},

	Route{
		Url:          "/admin/profile-pic",
		Method:       http.MethodPost,
		Handler:      controllers.AdminUploadImage,
		AuthRequired: true,
	},
}
