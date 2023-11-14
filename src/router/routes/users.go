package routes

import (
	"github.com/edigar/socialnets-web/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:                    "/criar-usuario",
		Method:                 http.MethodGet,
		Function:               controllers.CreateUser,
		AuthenticationRequired: false,
	},
	{
		URI:                    "/usuario",
		Method:                 http.MethodPost,
		Function:               controllers.StoreUser,
		AuthenticationRequired: false,
	},
	{
		URI:                    "/buscar-usuarios",
		Method:                 http.MethodGet,
		Function:               controllers.ShowUsers,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/usuario/{userId}",
		Method:                 http.MethodGet,
		Function:               controllers.ShowProfile,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/usuario/{userId}/parar-de-seguir",
		Method:                 http.MethodPost,
		Function:               controllers.UnfollowUser,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/usuario/{userId}/seguir",
		Method:                 http.MethodPost,
		Function:               controllers.FollowUser,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/perfil",
		Method:                 http.MethodGet,
		Function:               controllers.ShowLoggedUserProfile,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/editar-usuario",
		Method:                 http.MethodGet,
		Function:               controllers.ShowEditUser,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/editar-usuario",
		Method:                 http.MethodPut,
		Function:               controllers.EditUser,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/atualizar-senha",
		Method:                 http.MethodGet,
		Function:               controllers.ShowUpdatePassword,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/atualizar-senha",
		Method:                 http.MethodPost,
		Function:               controllers.UpdatePassword,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/deletar-usuario",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteUser,
		AuthenticationRequired: true,
	},
}
