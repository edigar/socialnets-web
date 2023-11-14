package routes

import (
	"github.com/edigar/socialnets-web/src/controllers"
	"net/http"
)

var homeRoute = Route{
	URI:                    "/home",
	Method:                 http.MethodGet,
	Function:               controllers.ShowHome,
	AuthenticationRequired: true,
}
