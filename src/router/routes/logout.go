package routes

import (
	"github.com/edigar/socialnets-web/src/controllers"
	"net/http"
)

var logoutRoute = Route{
	URI:                    "/logout",
	Method:                 http.MethodGet,
	Function:               controllers.Logout,
	AuthenticationRequired: true,
}
