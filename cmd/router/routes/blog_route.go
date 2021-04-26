package routes

import (
	"net/http"

	"github.com/scalent-sushil/user-management-go/pkg/handler"
)

var BlogsRoutes = []Route{

	Route{
		Url:          "/user/blogcreate",
		Method:       http.MethodPost,
		Handler:      handler.CreateBlog,
		AuthRequired: true,
	},

	Route{
		Url:          "/user/blogs",
		Method:       http.MethodGet,
		Handler:      handler.GetBlogs,
		AuthRequired: true,
	},

	Route{
		Url:          "/user/blog/{id}",
		Method:       http.MethodGet,
		Handler:      handler.GetBlog,
		AuthRequired: true,
	},

	Route{
		Url:          "/user/blog/{id}",
		Method:       http.MethodPut,
		Handler:      handler.Updateblog,
		AuthRequired: true,
	},

	Route{
		Url:          "/user/blog/{id}",
		Method:       http.MethodDelete,
		Handler:      handler.Deleteblog,
		AuthRequired: true,
	},
}
