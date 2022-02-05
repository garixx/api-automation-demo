package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

// Define our struct
type authenticationMiddleware struct {
	tokenUsers map[string]string
}

// Initialize it somewhere
func (amw *authenticationMiddleware) Populate() {
	if amw.tokenUsers == nil {
		amw.tokenUsers = make(map[string]string)
	}
	amw.tokenUsers["00000000"] = "user0"
	amw.tokenUsers["aaaaaaaa"] = "userA"
	amw.tokenUsers["05f717e5"] = "randomUser"
	amw.tokenUsers["deadbeef"] = "user0"
}

func (amw *authenticationMiddleware) Put(username string, token string){
	if amw.tokenUsers == nil {
		amw.tokenUsers = make(map[string]string)
	}

	amw.tokenUsers[token] = username
}

// Middleware function, which will be called for each request
func (amw *authenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if user, found := amw.tokenUsers[token]; found {
			// We found the token in our map
			logrus.Infof("Authenticated user %s\n", user)
			// Pass down the request to the next middleware (or final handler)
			next.ServeHTTP(w, r)
		} else {
			// Write an error and stop the handler chain
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}
