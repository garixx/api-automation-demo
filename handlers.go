package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

// LoginRequest authorize request payload
type LoginRequest struct {
	Username string     `json:"username,omitempty"`
	Password string     `json:"password,omitempty"`
	Token    string     `json:"token,omitempty"`
}

// struct for log response
// struct for logout response

type LogoutRequest struct {
	Token    string     `json:"token,omitempty"`
}

// Event represents data about a user event.
type event struct {
	Id          string   `json:"id,omitempty"`
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	Guests      []guest  `json:"guests,omitempty"`
	Time        string   `json:"time,omitempty"`
	Timezone    string   `json:"timezone,omitempty"`
	Duration    int      `json:"duration,omitempty"`
	Notes       []string `json:"notes,omitempty"`
}

type guest struct {
	Name       string    `json:"name,omitempty"`
	LastName   string    `json:"lastName,omitempty"`
	Company	   company   `json:"company,omitempty"`
}

type company struct {
	Name       string    `json:"name,omitempty"`
	Profile    string    `json:"profile,omitempty"`
}

var tokenUsers = map[string]string{
	"default": "default",
}

var exampleEvent1 = event{
	Id: "1",
	Title : "first event",
	Description: "sample event 1",
	Guests: []guest{
		{
			Name: "Peter",
			LastName: "North",
			Company: company{
				Name: "IBM",
				Profile: "IT Research" },
		},
		{
			Name: "John",
			LastName: "Dow",
		},
	},
	Time: "2021-12-29T19:03:19Z",
	Timezone: "UTC",
	Duration: 30,
	Notes: []string{"note1", "note2"},
}

var exampleEvent2 = event{
	Id: "2",
	Title : "second event",
	Description: "sample event 2",
	Time: "2021-12-24T07:25:19Z",
	Timezone: "Kyiv",
	Duration: 90,
	Notes: []string{"note3"},
}

var events = map[string]event{
	"1" : exampleEvent1,
	"2" : exampleEvent2,
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var auth LoginRequest

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&auth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	logrus.Info("Received login request: " + auth.Username + "," + auth.Password)

	newToken := base64.StdEncoding.EncodeToString([]byte(auth.Username + auth.Password))
	tokenUsers[newToken] = auth.Username

	logrus.Infof("tokens: %v", tokenUsers)
	fmt.Fprintf(w, newToken)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	var auth LoginRequest

	err := json.NewDecoder(r.Body).Decode(&auth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	logrus.Info("Received logout request")

	delete(tokenUsers, auth.Token)

	logrus.Infof("tokens: %v", tokenUsers)
	w.WriteHeader(http.StatusNoContent)
}

func DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	eventId := mux.Vars(r)["id"]
	logrus.Infof("Received delete event #%v request", eventId)

	delete(events, eventId)

	w.WriteHeader(http.StatusNoContent)
}

func EventsHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Received get events request")

	v := make([]event, 0, len(events))
	for  _, value := range events {
		v = append(v, value)
	}

	output, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func GetEventHandler(w http.ResponseWriter, r *http.Request) {
	eventId := mux.Vars(r)["id"]

	logrus.Infof("Received get event #%v request", eventId)

	output, err := json.Marshal(events[eventId])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)

	//// Read body
	//b, err := ioutil.ReadAll(r.Body)
	//defer r.Body.Close()
	//if err != nil {
	//	http.Error(w, err.Error(), 500)
	//	return
	//}
	//
	//// Unmarshal
	//var auth LoginRequest
	//err = json.Unmarshal(b, &auth)
	//if err != nil {
	//	http.Error(w, err.Error(), 500)
	//	return
	//}
	//
	//output, err := json.Marshal(auth)
	//if err != nil {
	//	http.Error(w, err.Error(), 500)
	//	return
	//}
	//w.Header().Set("content-type", "application/json")
	//w.Write(output)
}

func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("Received create new event request: %v", r.Body)
	var newEvent event

	err := json.NewDecoder(r.Body).Decode(&newEvent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	length := len(events) + 1
	newEventId := strconv.Itoa(length)

	newEvent.Id = newEventId
	newEvent.Time = time.Now().UTC().Format(time.RFC3339)

	events[newEventId] = newEvent

	output, err := json.Marshal(newEvent)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(output)
	//r.Context().
}

func PatchEventHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("Received patch new event request: %v", r.Body)
	var newEvent event

	err := json.NewDecoder(r.Body).Decode(&newEvent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//length := len(events) + 1
	//newEventId := strconv.Itoa(length)

	//newEvent.Id = newEventId
	newEvent.Time = time.Now().UTC().Format(time.RFC3339)

	oldEvent := events[newEvent.Id]

	//if newEvent.Title != nil {
	//	oldEvent.Title = newEvent.Title
	//}
	oldEvent.Title = newEvent.Title
	events[newEvent.Id] = oldEvent

	output, err := json.Marshal(newEvent)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(output)
}

//func AuthMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		var token string
//
//		if r.URL.Path == "/auth/login" || r.URL.Path == "/auth/logout" {
//			logrus.Info("authorization event")
//			token = "default"
//		} else {
//			logrus.Info("check is user authorized")
//			token = r.Header.Get("Authorization")
//		}
//
//		parsedToken := strings.Replace(token, "Bearer ", "", 1)
//
//		if user, found := tokenUsers[parsedToken]; found {
//			// We found the token in our map
//			logrus.Infof("Authenticated user: %s", user)
//			// Pass down the request to the next middleware (or final handler)
//			next.ServeHTTP(w, r)
//		} else {
//			// Write an error and stop the handler chain
//			http.Error(w, "Forbidden", http.StatusForbidden)
//		}
//	})
//}
