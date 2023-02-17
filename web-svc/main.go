package main

import (
	"flag"
	"fmt"
	"log"
	"max-web-svc/config"
	"max-web-svc/controller"
	"max-web-svc/repo"
	"max-web-svc/routes"
	"max-web-svc/service"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	controllers controller.Controllers
	svcs        service.Services
	repos       repo.Repositories
)

func main() {
	flag.Parse()

	router := mux.NewRouter()
	router.Host(config.Conf.Host)

	initRepos()
	initServices()
	initControllers()
	setupRoutes(router)
	handleIncomingRequests(router)
}

func initRepos() {
	repo.InitRepos(&repos)
}

func initServices() {
	svcs.Init(&repos)
}

func initControllers() {
	controllers.Init(&svcs)
}

func setupRoutes(apiSubrouter *mux.Router) {
	routes.SetupRoutes(apiSubrouter, &controllers)
}

func handleIncomingRequests(router *mux.Router) {
	done := make(chan error, 1)
	go func() {
		done <- http.ListenAndServe(fmt.Sprint(":", config.Conf.API_port), router)
	}()
	http.Handle("/", router)
	log.Printf("--- WEB SERVICE STARTED ---")
	if err := <-done; err != nil {
		log.Printf("Failed to serve. Error message: %s", err)
		os.Exit(1)
	}
}

func init() {
	//load default configuration
	if err := config.LoadConfJson(); err != nil {
		log.Fatalf("Failed loading service config. Error message: %s", err)
	}
	// load environment variables
	config.LoadEnv()
}
