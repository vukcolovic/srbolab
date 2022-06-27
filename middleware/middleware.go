package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"srbolabApp/handlers"
	"strings"
	"time"
)

func AuthToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "login") {
			next.ServeHTTP(w, r)
			return
		}

		cookie, err := r.Cookie("jwt")
		if err != nil {
			handlers.SetErrorResponse(w, errors.New("No cookie"))
			return
		}

		tknStr := cookie.Value
		claims := &jwt.StandardClaims{}

		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			c := &http.Cookie{
				Name:    "token",
				Value:   "",
				Path:    "/",
				Expires: time.Unix(0, 0),
				MaxAge:  -1,
			}

			http.SetCookie(w, c)
			handlers.SetErrorResponse(w, errors.New("No token"))
			return
		}
		if !tkn.Valid {
			c := &http.Cookie{
				Name:    "token",
				Value:   "",
				Path:    "/",
				Expires: time.Unix(0, 0),
				MaxAge:  -1,
			}

			http.SetCookie(w, c)
			handlers.SetErrorResponse(w, errors.New("Token not valid"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8080")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		//w.Header().Add("Access-Control-Allow-Headers", "X-Requested-With, X-HTTP-Method-Override, Authorization, Content-Type, Accept")
		//w.Header().Set("content-type", "application/json;charset=UTF-8")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
