package server

import (
	"../../data"
	u "../utils"
	"context"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
)
/*
Промежуточный слой аутентификации
 */

func JwtAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		notAuthPaths := NotAuthAccess
		requestPath := r.URL.Path
		for _, path := range notAuthPaths {
			if path == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}
		response := make(map[string] interface{})
		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			response = u.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}
		splitToken := strings.Split(tokenHeader, " ")
		if len(splitToken) != 2 {
			response = u.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}
		tokenPart := splitToken[1]
		tk := &data.AccessToken{}
		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})
		if err != nil {
			response = u.Message(false, "Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}
		if !token.Valid {
			response = u.Message(false, "Token is not valid")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}
		ctx := context.WithValue(r.Context(), "user", tk.UserUid)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}