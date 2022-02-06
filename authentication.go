package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

// data store
type authenticationMiddleware struct {
	tokenUsers map[string]string
}

// add token to token's store
func (amw *authenticationMiddleware) Put(username string, token string) {
	if amw.tokenUsers == nil {
		amw.tokenUsers = make(map[string]string)
	}

	amw.tokenUsers[token] = "Bearer " + username
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

		if user, found := tokenUsers[token]; found {
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
	for i := range list {
		if list[i] == str {
			return true
		}
	}
	return false
}
