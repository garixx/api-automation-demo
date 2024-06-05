package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"strconv"
)

const (
	swaggerRoute     = "/swagger"
	versionRoute     = "/version"
	loginRoute       = "/auth/login"
	logoutRoute      = "/auth/logout"
	createEventRoute = "/api/event"
	getEventsRoute   = "/api/events"
	getEventRoute    = "/api/event/{id:[0-9]+}"
	updateEventRoute = "/api/event/{id:[0-9]+}"
	deleteEventRoute = "/api/event/{id:[0-9]+}"
)

var routes = []string{loginRoute, logoutRoute, swaggerRoute}
var swaggerUrl string

func main() {
	swaggerUrl = *flag.String("swagger-url", "http://127.0.0.1:8087", "swagger docs url")
	url := flag.String("url", "http://127.0.0.1", "server port")
	port := flag.Int("port", 8081, "server port")
	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc(swaggerRoute, SwaggerHandler).Methods("GET")
	r.HandleFunc(versionRoute, VersionHandler).Methods("GET")
	r.HandleFunc(loginRoute, LoginHandler).Methods("POST")
	r.HandleFunc(logoutRoute, LogoutHandler).Methods("POST")
	r.HandleFunc(getEventsRoute, EventsHandler).Methods("GET")
	r.HandleFunc(getEventRoute, GetEventHandler).Methods("GET")
	r.HandleFunc(deleteEventRoute, DeleteEventHandler).Methods("DELETE")
	r.HandleFunc(createEventRoute, CreateEventHandler).Methods("POST")
	r.HandleFunc(updateEventRoute, UpdateEventHandler).Methods("PUT")

	r.Use(AuthMiddleware)

	logrus.Infof("Swagger port      : GET    %s:8087", *url)
	logrus.Infof("Authorize new user: POST %s. Payload: {\"username\": \"youruser\", \"password\": \"yourpass\"}", loginRoute)
	logrus.Infof("Drop user session : POST %s. Payload: {\"token\", \"yourtoken\"}", logoutRoute)
	logrus.Info("Get version        : GET    ", versionRoute)
	logrus.Info("Get all events     : GET    ", getEventsRoute)
	logrus.Info("Get event by id    : GET    ", getEventRoute)
	logrus.Info("Create event       : POST   ", createEventRoute)
	logrus.Info("Change event       : PUT    ", updateEventRoute)
	logrus.Info("Delete event       : DELETE ", deleteEventRoute)

	log.Println("Listening URL: ", *url)
	log.Println("Listening on port: ", *port)

	http.ListenAndServe(":"+strconv.Itoa(*port), r)
}
