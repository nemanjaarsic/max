package routes

import (
	"max-web-svc/controller"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router, ctrl *controller.Controllers) {
	router := r.PathPrefix("/api").Subrouter()

	router.HandleFunc("/login", ctrl.AuthController.Login).Methods("POST")
	router.HandleFunc("/deposit", ctrl.OperationController.Deposit).Methods("POST")
	router.HandleFunc("/withdraw", ctrl.OperationController.Withdraw).Methods("POST")
}
