package routes

import (
	"github.com/edigar/socialnets-web/src/controllers"
	"net/http"
)

var loginRoute = []Route{
	{
		URI:                    "/",
		Method:                 http.MethodGet,
		Function:               controllers.ShowLogin,
		AuthenticationRequired: false,
	},
	{
		URI:                    "/login",
		Method:                 http.MethodGet,
		Function:               controllers.ShowLogin,
		AuthenticationRequired: false,
	},
	{
		URI:                    "/login",
		Method:                 http.MethodPost,
		Function:               controllers.Login,
		AuthenticationRequired: false,
	},
}
