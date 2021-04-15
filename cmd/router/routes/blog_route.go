package routes

import (
	"net/http"

	controllers "github.com/scalent-sushil/user-management-go/pkg/handler"
)

var BlogsRoutes = []Route{

	Route{
		Url:          "/user/blogcreate",
		Method:       http.MethodPost,
		Handler:      controllers.CreateBlog,
		AuthRequired: true,
	},

	Route{
		Url:          "/user/blogs",
		Method:       http.MethodGet,
		Handler:      controllers.GetBlogs,
		AuthRequired: true,
	},

	Route{
		Url:          "/user/blog/{id}",
		Method:       http.MethodGet,
		Handler:      controllers.GetBlog,
		AuthRequired: true,
	},

	Route{
		Url:          "/user/blog/{id}",
		Method:       http.MethodPut,
		Handler:      controllers.Updateblog,
		AuthRequired: true,
	},

	Route{
		Url:          "/user/blog/{id}",
		Method:       http.MethodDelete,
		Handler:      controllers.Deleteblog,
		AuthRequired: true,
	},
}
