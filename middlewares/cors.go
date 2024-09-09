package middlewares

import (
	"net/http"
	"fmt"
)

func CorsMiddlewares(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
        fmt.Println("Origin received:", origin)

		allowedOrigins := map[string]bool{
			"https://fe-tb-berkah-jaya-750892348569.us-central1.run.app": true,
			"http://localhost:3000": true,
		}
		fmt.Println("Access-Control-Allow-Origin before:", w.Header().Get("Access-Control-Allow-Origin"))

		if allowedOrigins[origin] {
			if w.Header().Get("Access-Control-Allow-Origin") == "" {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Method", "GET, POST, PUT, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Header", "X-Requested-With, Content-Type, Authorization")
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}
		}

		fmt.Println("Access-Control-Allow-Origin after:", w.Header().Get("Access-Control-Allow-Origin"))
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}