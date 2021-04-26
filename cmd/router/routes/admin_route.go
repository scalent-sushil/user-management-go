package routes

import (
	"net/http"

	"github.com/scalent-sushil/user-management-go/pkg/handler"
)

var AdminRoutes = []Route{

	Route{
		Url:          "/admin/users",
		Method:       http.MethodGet,
		Handler:      handler.AdminGetUsers,
		AuthRequired: true,
	},
	Route{
		Url:          "/admin",
		Method:       http.MethodGet,
		Handler:      handler.AdminProfile,
		AuthRequired: true,
	},
	Route{
		Url:          "/admin/user/{id}",
		Method:       http.MethodGet,
		Handler:      handler.GetUser,
		AuthRequired: true,
	},
	Route{
		Url:          "/admin/user/{id}",
		Method:       http.MethodPut,
		Handler:      handler.AdminUpdateUser,
		AuthRequired: true,
	},

	Route{
		Url:          "/admin/user/{id}/status",
		Method:       http.MethodPut,
		Handler:      handler.DeleteUserByAdmin,
		AuthRequired: true,
	},

	Route{
		Url:          "/admin/profile-pic",
		Method:       http.MethodPost,
		Handler:      handler.AdminUploadImage,
		AuthRequired: true,
	},
}
