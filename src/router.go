package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func NewRouter(repo Repository) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	var routes = defineRoutes(&UserService{repo})
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(handler)
	}

	return router
}

func defineRoutes(service *UserService) Routes {
	return Routes{
		Route{
			"Index",
			"GET",
			"/",
			service.Index,
		},
		Route{
			"UserList",
			"GET",
			"/users",
			service.GetUserList,
		},
		Route{
			"CreateUser",
			"POST",
			"/users",
			service.CreateUser,
		},
		Route{
			"UserProfile",
			"GET",
			"/users/{userId}",
			service.GetUserProfile,
		},
		Route{
			"SignUpUser",
			"POST",
			"/signup",
			service.SignUpUser,
		},
	}
}
