package handlers

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
	"os"
)

func AuthHandler(h http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if ok {
			username := sha256.Sum256([]byte(os.Getenv("USER_NAME")))
			password := sha256.Sum256([]byte(os.Getenv("USER_PASS")))
			userHash := sha256.Sum256([]byte(user))
			passHash := sha256.Sum256([]byte(pass))
			validUser := subtle.ConstantTimeCompare(userHash[:], username[:]) == 1
			validPass := subtle.ConstantTimeCompare(passHash[:], password[:]) == 1
			if validPass && validUser {
				h.ServeHTTP(rw, r)
				return
			}
		}
		http.Error(rw, "No/Invalid Credentials", http.StatusUnauthorized)
	}
}
