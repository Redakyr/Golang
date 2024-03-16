package routes

import (
	"github.com/Redakyr/Golang/GO_API/pkg/controllers/"
	"github.com/gorilla/mux"
)

var registerAgentRoutes = func(router *mux.Router) {
	router.HandleFunc("/whistleblower/", controllers.RegisterAgent).Methods("POST")
	router.HandleFunc("/whistleblower/", controllers.GetAgent).Methods("GET")
	router.HandleFunc("/whistleblower/{agentId}", controllers.GetAgentById).Methods("GET")
	router.HandleFunc("/whistleblower/{agentId}", controllers.UpdateAgent).Methods("PUT")
	router.HandleFunc("/whistleblower/{agentId}", controllers.RemoveAgent).Methods("DELETE")
}
