package middlewares

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"golang.org/x/net/http2"
	helper "github.com/RaihanMalay21/helper_TB_Berkah_Jaya"
)

func ReverseProxy(target string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		url, err := url.Parse(target)
		if err != nil {
			log.Println("Error Parsing URL: ", err)
			message := map[string]interface{}{"message": err}
			helper.Response(w, message, http.StatusInternalServerError)
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(url)

        // Set custom Transport to support HTTP/2
        proxy.Transport = &http2.Transport{}

		pathMicroservices := []string{
			"https://server-customer-tb-berkah-jaya-750892348569.us-central1.run.app",
			"https://server-registry-tb-berkah-jaya-750892348569.us-central1.run.app",
		}

		PathPrefix := []string{
			"/customer",
			"/access",
		}

		for i, path := range pathMicroservices {
			if path == url.String() {
				r.URL.Path = strings.TrimPrefix(r.URL.Path, PathPrefix[i])
				break
			}
		}
		
		log.Println("Proxying to:", url.String())
		log.Println("Request URL:", r.URL.String())
		log.Println("Request Host:", r.Host)
		log.Println("Request Scheme:", r.URL.Scheme)

        // Modify the request
        r.URL.Host = url.Host
        r.URL.Scheme = url.Scheme
        r.Header.Set("X-Forwarded-Host", r.Host)
        r.Host = url.Host

		proxy.ServeHTTP(w, r)
	})
}