package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	loginRoute = "/auth/login"
	logoutRoute = "/auth/logout"
	createEventRoute = "/api/event"
	getEventsRoute = "/api/events"
	getEventRoute = "/api/event/{id:[0-9]+}"
	deleteEventRoute = "/api/event/{id:[0-9]+}"

)

var routes = []string {loginRoute, logoutRoute}

func main() {
	port := flag.Int("port", 8081, "server port")
	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc(loginRoute, LoginHandler).Methods("POST")
	r.HandleFunc(logoutRoute, LogoutHandler).Methods("POST")
	r.HandleFunc(getEventsRoute, EventsHandler).Methods("GET")
	r.HandleFunc(getEventRoute, GetEventHandler).Methods("GET")
	r.HandleFunc(deleteEventRoute, DeleteEventHandler).Methods("DELETE")
	r.HandleFunc(createEventRoute,CreateEventHandler).Methods("POST")

	r.Use(AuthMiddleware)

	logrus.Infof("Authorize new user: POST %s", loginRoute)
	logrus.Info("Drop user session : POST %s", logoutRoute)
	logrus.Info("Get all events    : GET %s", getEventsRoute)
	logrus.Info("Get event by id   : GET %s", getEventRoute)
	logrus.Info("Create event      : POST %s", createEventRoute)
	logrus.Info("Delete event      : DELETE %s", deleteEventRoute)

	log.Printf("Listening URL: http://localhost")
	log.Println("Listening on port: ", *port)

	http.ListenAndServe(":"+strconv.Itoa(*port), r)
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var token string

		if contains(routes, r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}

		logrus.Info("check is user authorized")
		token = r.Header.Get("Authorization")

		parsedToken := strings.Replace(token, "Bearer ", "", 1)

		if user, found := tokenUsers[parsedToken]; found {
			// We found the token in our map
			logrus.Infof("Authenticated user: %s", user)
			// Pass down the request to the next middleware (or final handler)
			next.ServeHTTP(w, r)
		} else {
			// Write an error and stop the handler chain
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

func contains(list []string, str string) bool {
	for i :=  range list {
		if list[i] == str {
			return true
		}
	}
	return false
}
