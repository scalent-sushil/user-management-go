package routes

import (
	"net/http"

	"github.com/scalent-sushil/user-management-go/pkg/handler"
)

var LoginRoutes = []Route{
	Route{
		Url:          "/login",
		Method:       http.MethodPost,
		Handler:      handler.Login,
		AuthRequired: false,
	},
	Route{
		Url:          "/adminlogin",
		Method:       http.MethodPost,
		Handler:      handler.AdminLogin,
		AuthRequired: false,
	},
	Route{
		Url:          "/registration",
		Method:       http.MethodPost,
		Handler:      handler.Registration,
		AuthRequired: false,
	},
	Route{
		Url:          "/forgotpassword",
		Method:       http.MethodPost,
		Handler:      handler.ForgotPassword,
		AuthRequired: false,
	},
	Route{
		Url:          "/newpassword",
		Method:       http.MethodPost,
		Handler:      handler.NewPassword,
		AuthRequired: false,
	},
}
