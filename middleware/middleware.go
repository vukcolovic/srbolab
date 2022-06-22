package middleware

import (
	"fmt"
	"net/http"
)

func AuthToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			fmt.Println(cookie)
			return
		}
		next.ServeHTTP(w, r)
	})
}
