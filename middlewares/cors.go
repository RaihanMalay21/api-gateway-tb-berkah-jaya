package middlewares

import (
	"net/http"
	"fmt"
)

func CorsMiddlewares(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
        fmt.Println("Origin received:", origin)

		allowedOrigins := []string{
			"https://fe-tb-berkah-jaya-750892348569.us-central1.run.app",
			"http://localhost:3000",
		}
		fmt.Println("Access-Control-Allow-Origin before:", w.Header().Get("Access-Control-Allow-Origin"))
		for _, Allow := range allowedOrigins {
			if Allow == origin {
				w.Header().Set("Access-Control-Allow-Origin", Allow)
				w.Header().Set("Access-Control-Allow-Method", "GET, POST, PUT, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Header", "X-Requested-With, Content-Type, Authorization")
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				break
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