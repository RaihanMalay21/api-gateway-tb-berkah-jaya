package middlewares

import (
	"net/http"
)

func CorsMiddlewares(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		allowedOrigins := []string{
			"https://fe-tb-berkah-jaya-750892348569.us-central1.run.app",
		}

		for _, Allow := range allowedOrigins {
			if Allow == origin {
				w.Header().Set("Access-Control-Allow-Origin", Allow)
				w.Header().Set("Access-Control-Allow-Method", "GET, POST, PUT, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Header", "X-Requested-With, Content-Type, Authorization")
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				break
			}
		}

		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}