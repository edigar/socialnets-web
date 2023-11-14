package routes

import (
	"github.com/edigar/socialnets-web/src/controllers"
	"net/http"
)

var postRoutes = []Route{
	{
		URI:                    "/post",
		Method:                 http.MethodPost,
		Function:               controllers.PostPost,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/post",
		Method:                 http.MethodGet,
		Function:               controllers.GetPosts,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/post/{postId}/edit",
		Method:                 http.MethodGet,
		Function:               controllers.EditPosts,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/post/{postId}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdatePost,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/post/{postId}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeletePost,
		AuthenticationRequired: true,
	},
	//{
	//	URI:                    "/user/{userId}/posts",
	//	Method:                 http.MethodGet,
	//	Function:               controllers.GetUserPosts,
	//	AuthenticationRequired: true,
	//},
	{
		URI:                    "/post/{postId}/like",
		Method:                 http.MethodPost,
		Function:               controllers.LikePost,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/post/{postId}/unlike",
		Method:                 http.MethodPost,
		Function:               controllers.UnlikePost,
		AuthenticationRequired: true,
	},
}
