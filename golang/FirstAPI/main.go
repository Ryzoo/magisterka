package main

import (
	"FirstAPI/pkg/Dependencies"
	"FirstAPI/pkg/Users"
	"net/http"
)

func main() {
	dependency := Dependencies.InitDependencies()

	registerControllers(dependency)
	startServer()
}

func startServer() {
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		println(err.Error())
	}
}

func registerControllers(dependencies Dependencies.Dependencies){
	http.HandleFunc("/users", Users.GetNewUsersController(dependencies.Database.UserRepository).Handle)
}

